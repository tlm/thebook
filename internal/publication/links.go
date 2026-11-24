package publication

import (
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

// localLink is a destination resolved to a path within the published site.
type localLink struct {
	fragment string
	query    string
	target   string
}

type markdownDocument struct {
	anchors map[string]bool
	links   []string
}

// markdown parses actual links and heading IDs, excluding code examples.
// Raw HTML links and external destinations are outside this validator.
func markdown(data []byte) markdownDocument {
	document := parseMarkdown(data)
	result := markdownDocument{anchors: make(map[string]bool)}
	collectNode := func(node ast.Node, entering bool) (
		ast.WalkStatus,
		error,
	) {
		if !entering {
			return ast.WalkContinue, nil
		}
		switch node := node.(type) {
		case *ast.Heading:
			id, ok := node.AttributeString("id")
			if ok {
				result.anchors[string(id.([]byte))] = true
			}
		case *ast.Image:
			result.links = append(
				result.links,
				string(node.Destination),
			)
		case *ast.Link:
			result.links = append(
				result.links,
				string(node.Destination),
			)
		}
		return ast.WalkContinue, nil
	}
	ast.Walk(document, collectNode)
	return result
}

// parseMarkdown parses data with the heading IDs that link anchors match.
func parseMarkdown(data []byte) ast.Node {
	engine := goldmark.New(goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	))
	return engine.Parser().Parse(text.NewReader(data))
}

// resolveLink resolves a Markdown destination against the published path of
// the document containing it, reporting false for external destinations.
// It rejects site-root paths because project Pages sites live under a
// repository prefix. Queries do not change the target file.
func resolveLink(from string, destination string) (localLink, bool, error) {
	// Match Markdown rendering of punctuation escapes and HTML entities.
	escaped := util.URLEscape([]byte(destination), true)
	link, err := url.Parse(string(escaped))
	if err != nil {
		return localLink{}, false, fmt.Errorf(
			"parsing destination %q: %w",
			destination,
			err,
		)
	}
	if link.IsAbs() || link.Host != "" {
		return localLink{}, false, nil
	}
	if strings.Contains(link.Path, "\\") ||
		strings.HasPrefix(link.Path, "/") {
		return localLink{}, false, fmt.Errorf(
			"non-relative destination %q",
			destination,
		)
	}
	target := from
	if link.Path != "" {
		target = path.Join(path.Dir(from), link.Path)
	}
	if target == ".." || strings.HasPrefix(target, "../") {
		return localLink{}, false, fmt.Errorf(
			"destination escapes site: %q",
			destination,
		)
	}
	resolved := localLink{
		fragment: link.Fragment,
		query:    link.RawQuery,
		target:   target,
	}
	return resolved, true, nil
}

// validateLinks checks Markdown destinations against the assembled snapshot,
// not the source tree, so excluded files cannot satisfy a published link.
func validateLinks(files map[string][]byte) error {
	documents := make(map[string]markdownDocument)
	for name, data := range files {
		if strings.HasSuffix(name, ".md") {
			documents[name] = markdown(data)
		}
	}
	for name, document := range documents {
		for _, destination := range document.links {
			err := validateTarget(
				name,
				destination,
				files,
				documents,
			)
			if err != nil {
				return fmt.Errorf("link in %q: %w", name, err)
			}
		}
	}
	return nil
}

// validateTarget checks that a local destination names a published file
// and, for Markdown files, an existing heading.
func validateTarget(
	from string,
	destination string,
	files map[string][]byte,
	documents map[string]markdownDocument,
) error {
	link, local, err := resolveLink(from, destination)
	if err != nil {
		return err
	}
	if !local {
		return nil
	}
	if _, ok := files[link.target]; !ok {
		return fmt.Errorf("unpublished destination %q", destination)
	}
	if link.fragment != "" && strings.HasSuffix(link.target, ".md") {
		if !documents[link.target].anchors[link.fragment] {
			return fmt.Errorf("missing heading in %q", destination)
		}
	}
	return nil
}
