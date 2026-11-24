package publication

import (
	"bytes"
	"cmp"
	"fmt"
	"maps"
	"net/url"
	"path"
	"slices"
	"strings"

	"github.com/yuin/goldmark/ast"
)

// destinationEdit replaces the bytes of one link destination.
type destinationEdit struct {
	replacement string
	start       int
	stop        int
}

// headingPosition locates a heading and the line that starts its section.
type headingPosition struct {
	id       string
	level    int
	start    int
	topLevel bool
}

// linkPosition locates a parsed link or image within its document.
type linkPosition struct {
	destination string
	reference   bool
	start       int
}

// referencePosition locates a link reference definition.
type referencePosition struct {
	destination string
	start       int
	stop        int
}

// sectionKey identifies the section a heading starts in a source document.
type sectionKey struct {
	anchor   string
	document string
}

// sectionPublisher extracts linked heading sections into standalone files
// and rewrites links to them. Sections are extracted from the original
// source bytes, so the result does not depend on the order of rewriting.
type sectionPublisher struct {
	files    map[string][]byte
	sections map[sectionKey]string
	sources  map[string][]byte
}

const (
	// destinationEnds lists bytes that can follow an unbracketed
	// destination.
	destinationEnds = ") \t\r\n"

	// sectionRoot is the reserved root directory for published sections.
	sectionRoot = "sections"

	// whitespace lists bytes that can precede a destination.
	whitespace = " \t\r\n"
)

// applyEdits returns a copy of data with each edit applied.
func applyEdits(data []byte, edits []destinationEdit) ([]byte, error) {
	byStart := func(a, b destinationEdit) int {
		return cmp.Compare(a.start, b.start)
	}
	slices.SortFunc(edits, byStart)
	result := make([]byte, 0, len(data))
	offset := 0
	for _, edit := range edits {
		if edit.start < offset {
			return nil, fmt.Errorf(
				"overlapping destination edits at offset %d",
				edit.start,
			)
		}
		result = append(result, data[offset:edit.start]...)
		result = append(result, edit.replacement...)
		offset = edit.stop
	}
	return append(result, data[offset:]...), nil
}

// destination returns the replacement for raw, a destination written
// relative to origin, when it appears in the published file document.
// local maps anchors in origin to their IDs in document when document is
// an extracted section of origin.
func (p *sectionPublisher) destination(
	document string,
	origin string,
	raw string,
	local map[string]string,
) (string, error) {
	link, isLocal, err := resolveLink(origin, raw)
	if err != nil {
		return "", err
	}
	if !isLocal {
		return raw, nil
	}
	_, isSource := p.sources[link.target]
	extract := link.fragment != "" && isSource &&
		strings.HasSuffix(link.target, ".md")
	if extract && link.target == origin {
		if document == origin {
			return raw, nil
		}
		id, inside := local[link.fragment]
		if inside {
			sameSection := url.URL{Fragment: id}
			return sameSection.String(), nil
		}
	}
	if extract {
		var name string
		name, err = p.section(link.target, link.fragment)
		if err != nil {
			return "", err
		}
		return relativeDestination(
			path.Dir(document),
			name,
			link.query,
			"",
		), nil
	}
	if document == origin {
		return raw, nil
	}
	return relativeDestination(
		path.Dir(document),
		link.target,
		link.query,
		link.fragment,
	), nil
}

// findDestination returns the offset of destination where it follows
// marker, optional whitespace and an optional opening angle bracket within
// data[from:stop]. The destination must end where its form ends it.
func findDestination(
	data []byte,
	from int,
	stop int,
	marker string,
	destination string,
) (int, bool) {
	for offset := from; offset < stop; {
		index := bytes.Index(data[offset:stop], []byte(marker))
		if index < 0 {
			return 0, false
		}
		start := offset + index + len(marker)
		offset = start
		for start < stop &&
			strings.IndexByte(whitespace, data[start]) >= 0 {
			start++
		}
		angled := start < stop && data[start] == '<'
		if angled {
			start++
		}
		end := start + len(destination)
		if end > stop || string(data[start:end]) != destination {
			continue
		}
		ended := end == stop ||
			strings.IndexByte(destinationEnds, data[end]) >= 0
		if angled && end < stop && data[end] == '>' {
			return start, true
		} else if !angled && ended {
			return start, true
		}
	}
	return 0, false
}

// headingPositions lists every heading in document order.
func headingPositions(document ast.Node, data []byte) []headingPosition {
	var headings []headingPosition
	collectHeading := func(node ast.Node, entering bool) (
		ast.WalkStatus,
		error,
	) {
		heading, ok := node.(*ast.Heading)
		if !entering || !ok {
			return ast.WalkContinue, nil
		}
		var id string
		value, ok := heading.AttributeString("id")
		if ok {
			id = string(value.([]byte))
		}
		before := data[:max(heading.Pos(), 0)]
		headings = append(headings, headingPosition{
			id:       id,
			level:    heading.Level,
			start:    bytes.LastIndexByte(before, '\n') + 1,
			topLevel: heading.Parent().Kind() == ast.KindDocument,
		})
		return ast.WalkContinue, nil
	}
	ast.Walk(document, collectHeading)
	return headings
}

// linkPositions lists links and images in document order, along with the
// link reference definitions that reference links take destinations from.
func linkPositions(
	document ast.Node,
) ([]linkPosition, []referencePosition) {
	var links []linkPosition
	var references []referencePosition
	collectNode := func(node ast.Node, entering bool) (
		ast.WalkStatus,
		error,
	) {
		if !entering {
			return ast.WalkContinue, nil
		}
		switch node := node.(type) {
		case *ast.Image:
			links = append(links, linkPosition{
				destination: string(node.Destination),
				reference:   node.Reference != nil,
				start:       node.Pos(),
			})
		case *ast.Link:
			links = append(links, linkPosition{
				destination: string(node.Destination),
				reference:   node.Reference != nil,
				start:       node.Pos(),
			})
		case *ast.LinkReferenceDefinition:
			lines := node.Lines()
			if lines.Len() > 0 {
				references = append(references, referencePosition{
					destination: string(node.Destination),
					start:       lines.At(0).Start,
					stop:        lines.At(lines.Len() - 1).Stop,
				})
			}
		}
		return ast.WalkContinue, nil
	}
	ast.Walk(document, collectNode)
	return links, references
}

// publishSections gives each heading that one Markdown file links to in
// another its own published file under [sectionRoot], holding that
// heading's section, and rewrites those links to it. Links within a
// section are rebased to its location. Same-document links, links without
// a fragment and external links are unchanged. It adds section files and
// replaces rewritten documents in files.
func publishSections(files map[string][]byte) error {
	publisher := sectionPublisher{
		files:    files,
		sections: make(map[sectionKey]string),
		sources:  maps.Clone(files),
	}
	for _, name := range slices.Sorted(maps.Keys(publisher.sources)) {
		if !strings.HasSuffix(name, ".md") {
			continue
		}
		data, err := publisher.rewriteLinks(
			name,
			name,
			publisher.sources[name],
			nil,
		)
		if err != nil {
			return fmt.Errorf("rewriting links in %q: %w", name, err)
		}
		files[name] = data
	}
	return nil
}

// relativeDestination returns a link destination from directory dir to
// target, both relative to the site root.
func relativeDestination(
	dir string,
	target string,
	query string,
	fragment string,
) string {
	destination := url.URL{
		Path:     relativePath(dir, target),
		RawQuery: query,
		Fragment: fragment,
	}
	return destination.String()
}

// relativePath returns the slash-separated path from directory dir to the
// file target, both relative to the site root.
func relativePath(dir string, target string) string {
	var from []string
	if dir != "." {
		from = strings.Split(dir, "/")
	}
	to := strings.Split(target, "/")
	common := 0
	for common < len(from) && common < len(to)-1 &&
		from[common] == to[common] {
		common++
	}
	parts := make([]string, 0, len(from)+len(to)-2*common)
	for range len(from) - common {
		parts = append(parts, "..")
	}
	parts = append(parts, to[common:]...)
	return strings.Join(parts, "/")
}

// rewriteLinks returns data, written relative to origin, with each link
// destination replaced as [sectionPublisher.destination] determines for
// the published file document. It parses the result again and fails
// unless every link changed exactly as intended, so a link form it cannot
// rewrite precisely is reported rather than published incorrectly.
func (p *sectionPublisher) rewriteLinks(
	document string,
	origin string,
	data []byte,
	local map[string]string,
) ([]byte, error) {
	links, references := linkPositions(parseMarkdown(data))
	replacements := make(map[string]string)
	for _, link := range links {
		_, resolved := replacements[link.destination]
		if resolved {
			continue
		}
		replacement, err := p.destination(
			document,
			origin,
			link.destination,
			local,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"resolving %q: %w",
				link.destination,
				err,
			)
		}
		replacements[link.destination] = replacement
	}

	edits := make([]destinationEdit, 0, len(links))
	for _, link := range links {
		replacement := replacements[link.destination]
		if link.reference || replacement == link.destination {
			continue
		}
		start, found := findDestination(
			data,
			link.start,
			len(data),
			"](",
			link.destination,
		)
		if !found {
			return nil, fmt.Errorf(
				"locating destination %q",
				link.destination,
			)
		}
		edits = append(edits, destinationEdit{
			replacement: replacement,
			start:       start,
			stop:        start + len(link.destination),
		})
	}
	for _, reference := range references {
		replacement, used := replacements[reference.destination]
		if !used || replacement == reference.destination {
			continue
		}
		start, found := findDestination(
			data,
			reference.start,
			reference.stop,
			"]:",
			reference.destination,
		)
		if !found {
			return nil, fmt.Errorf(
				"locating reference destination %q",
				reference.destination,
			)
		}
		edits = append(edits, destinationEdit{
			replacement: replacement,
			start:       start,
			stop:        start + len(reference.destination),
		})
	}
	if len(edits) == 0 {
		return data, nil
	}

	result, err := applyEdits(data, edits)
	if err != nil {
		return nil, err
	}
	want := make([]string, 0, len(links))
	for _, link := range links {
		want = append(want, replacements[link.destination])
	}
	if !slices.Equal(markdown(result).links, want) {
		return nil, fmt.Errorf(
			"rewritten links do not match their intended destinations",
		)
	}
	return result, nil
}

// section returns the published path of the section that anchor starts in
// document, extracting it on first use and reusing that path afterwards.
// A path already published for anything else indicates a publisher
// defect, so it returns an error rather than replacing that file.
func (p *sectionPublisher) section(
	document string,
	anchor string,
) (string, error) {
	key := sectionKey{anchor: anchor, document: document}
	name, extracted := p.sections[key]
	if extracted {
		return name, nil
	}
	content, local, err := sectionContent(p.sources[document], anchor)
	if err != nil {
		return "", fmt.Errorf(
			"extracting section %q of %q: %w",
			anchor,
			document,
			err,
		)
	}
	name = sectionPath(document, anchor)
	if _, exists := p.files[name]; exists {
		return "", fmt.Errorf(
			"section %q of %q would replace published file %q",
			anchor,
			document,
			name,
		)
	}
	// Reserve the path first so links leading back to this section
	// reuse it instead of extracting it again.
	p.sections[key] = name
	p.files[name] = nil
	content, err = p.rewriteLinks(name, document, content, local)
	if err != nil {
		return "", fmt.Errorf(
			"rewriting links in section %q of %q: %w",
			anchor,
			document,
			err,
		)
	}
	header := fmt.Sprintf(
		"> From [`%s`](%s).\n\n",
		document,
		relativeDestination(path.Dir(name), document, "", ""),
	)
	p.files[name] = append([]byte(header), content...)
	return name, nil
}

// sectionContent returns the section that the heading with anchor starts
// in data, ending before the next heading at the top level with the same
// or a higher rank. It also maps the source anchors of headings in the
// section to their IDs in the extracted content, which differ when a
// duplicate heading precedes the section.
func sectionContent(
	data []byte,
	anchor string,
) ([]byte, map[string]string, error) {
	document := parseMarkdown(data)
	headings := headingPositions(document, data)
	matchesAnchor := func(heading headingPosition) bool {
		return heading.id == anchor
	}
	index := slices.IndexFunc(headings, matchesAnchor)
	if index < 0 {
		return nil, nil, fmt.Errorf("missing heading %q", anchor)
	}
	heading := headings[index]
	if !heading.topLevel {
		return nil, nil, fmt.Errorf(
			"heading %q is nested in another block",
			anchor,
		)
	}
	stop := len(data)
	inner := headings[index:]
	for offset, next := range headings[index+1:] {
		if next.topLevel && next.level <= heading.level {
			stop = next.start
			inner = headings[index : index+1+offset]
			break
		}
	}
	content := bytes.Clone(data[heading.start:stop])

	extracted := headingPositions(parseMarkdown(content), content)
	if len(extracted) != len(inner) {
		return nil, nil, fmt.Errorf(
			"headings in section %q differ once extracted",
			anchor,
		)
	}
	sourceLinks, _ := linkPositions(document)
	inSection := 0
	for _, link := range sourceLinks {
		if link.start >= heading.start && link.start < stop {
			inSection++
		}
	}
	extractedLinks, _ := linkPositions(parseMarkdown(content))
	if len(extractedLinks) != inSection {
		return nil, nil, fmt.Errorf(
			"links in section %q depend on content outside it",
			anchor,
		)
	}
	anchors := make(map[string]string, len(inner))
	for position, inside := range inner {
		anchors[inside.id] = extracted[position].id
	}
	return content, anchors, nil
}

// sectionPath names the published file for the section that anchor starts
// in document. Heading IDs contain only lowercase letters, digits and
// hyphens, so distinct sections always receive distinct paths.
func sectionPath(document string, anchor string) string {
	return path.Join(
		sectionRoot,
		strings.TrimSuffix(document, ".md"),
		anchor+".md",
	)
}
