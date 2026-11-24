package publication

import (
	"maps"
	"slices"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/juju/tc"
)

// sectionsSuite exercises publication of linked heading sections.
type sectionsSuite struct{}

// publishedSections lists the published section files in a bundle.
func publishedSections(bundle Bundle) []string {
	var names []string
	for _, name := range slices.Sorted(maps.Keys(bundle.Files)) {
		if strings.HasPrefix(name, "sections/") {
			names = append(names, name)
		}
	}
	return names
}

// sectionSource provides a guide with nested sections for other files to
// link to.
func sectionSource() fstest.MapFS {
	return fstest.MapFS{
		"HARNESS.md": {
			Data: []byte("# Harness\n\n[Guide](go/guide.md)\n"),
		},
		"go/guide.md": {
			Data: []byte("# Go Guide\n\n" +
				"## Errors\n\nWrap errors.\n\n" +
				"### Matching\n\nUse errors.Is.\n\n" +
				"## Testing\n\nTest behaviour.\n"),
		},
		"go/details.md": {
			Data: []byte("# Details\n\nMore detail.\n"),
		},
		"assets/raw.bin": {Data: []byte{0, 255}},
	}
}

// TestSections registers the linked section publication contract tests.
func TestSections(t *testing.T) {
	tc.Run(t, &sectionsSuite{})
}

// TestExtractsLinkedSection verifies a linked heading's section, including
// its subsections, is published alone and the link is rewritten to it.
func (sectionsSuite) TestExtractsLinkedSection(c *tc.C) {
	source := sectionSource()
	source["HARNESS.md"].Data = []byte(
		"# Harness\n\n[Errors](go/guide.md#errors)\n",
	)
	bundle, err := Build(source)
	c.Assert(err, tc.IsNil)
	c.Check(
		string(bundle.Files["sections/go/guide/errors.md"]),
		tc.Equals,
		"> From [`go/guide.md`](../../../go/guide.md).\n\n"+
			"## Errors\n\nWrap errors.\n\n"+
			"### Matching\n\nUse errors.Is.\n\n",
	)
	c.Check(
		string(bundle.Files["HARNESS.md"]),
		tc.Equals,
		"# Harness\n\n[Errors](sections/go/guide/errors.md)\n",
	)
	c.Check(
		bundle.Files["go/guide.md"],
		tc.DeepEquals,
		source["go/guide.md"].Data,
	)
}

// TestKeepsOtherLinks verifies links without a fragment for another
// Markdown file leave every file unchanged and publish no sections.
func (sectionsSuite) TestKeepsOtherLinks(c *tc.C) {
	source := sectionSource()
	source["HARNESS.md"].Data = []byte("# Harness\n\n" +
		"[Guide](go/guide.md)\n" +
		"[Web](https://example.invalid/page#errors)\n" +
		"[Asset](assets/raw.bin)\n" +
		"[Here](#harness)\n")
	source["go/guide.md"].Data = []byte("# Go Guide\n\n" +
		"## Errors\n\n[Testing](#testing) and [Errors](guide.md#errors).\n\n" +
		"## Testing\n\nTest behaviour.\n")
	bundle, err := Build(source)
	c.Assert(err, tc.IsNil)
	for name, file := range source {
		c.Check(bundle.Files[name], tc.DeepEquals, file.Data)
	}
	c.Check(publishedSections(bundle), tc.HasLen, 0)
}

// TestRebasesSectionLinks verifies links inside a section resolve from its
// published location, keeping same-section anchors and publishing other
// linked sections in turn.
func (sectionsSuite) TestRebasesSectionLinks(c *tc.C) {
	source := sectionSource()
	source["HARNESS.md"].Data = []byte(
		"# Harness\n\n[Errors](go/guide.md#errors)\n",
	)
	source["go/guide.md"].Data = []byte("# Go Guide\n\n" +
		"## Errors\n\n" +
		"[Home](../HARNESS.md) ![Asset](../assets/raw.bin)\n" +
		"[Matching](#matching) [Testing](#testing)\n" +
		"[Details](details.md#details)\n\n" +
		"### Matching\n\nUse errors.Is.\n\n" +
		"## Testing\n\nTest behaviour.\n")
	bundle, err := Build(source)
	c.Assert(err, tc.IsNil)
	c.Check(
		string(bundle.Files["sections/go/guide/errors.md"]),
		tc.Equals,
		"> From [`go/guide.md`](../../../go/guide.md).\n\n"+
			"## Errors\n\n"+
			"[Home](../../../HARNESS.md) "+
			"![Asset](../../../assets/raw.bin)\n"+
			"[Matching](#matching) [Testing](testing.md)\n"+
			"[Details](../details/details.md)\n\n"+
			"### Matching\n\nUse errors.Is.\n\n",
	)
	c.Check(publishedSections(bundle), tc.DeepEquals, []string{
		"sections/go/details/details.md",
		"sections/go/guide/errors.md",
		"sections/go/guide/testing.md",
	})
}

// TestRejectsUnsupportedSections verifies sections without a clean
// boundary, sections that depend on outside content, and links that cannot
// be rewritten exactly cause errors rather than incorrect output.
func (sectionsSuite) TestRejectsUnsupportedSections(c *tc.C) {
	cases := []struct {
		guide   string
		harness string
		name    string
	}{
		{
			name:    "nested heading",
			harness: "[Quoted](go/guide.md#quoted)",
			guide:   "# Go Guide\n\n> ## Quoted\n>\n> Text.\n",
		},
		{
			name:    "outside reference definition",
			harness: "[Errors](go/guide.md#errors)",
			guide: "# Go Guide\n\n## Errors\n\n[Details][d]\n\n" +
				"## Testing\n\n[d]: details.md\n",
		},
		{
			name: "ambiguous destination",
			harness: "[![Errors](go/guide.md#errors)]" +
				"(go/guide.md#errors)",
			guide: "# Go Guide\n\n## Errors\n",
		},
	}
	for _, test := range cases {
		checkRejected := func(c *testing.T) {
			source := sectionSource()
			source["HARNESS.md"].Data = []byte(
				"# Harness\n\n" + test.harness + "\n",
			)
			source["go/guide.md"].Data = []byte(test.guide)
			_, err := Build(source)
			tc.Check(c, err == nil, tc.IsFalse)
		}
		c.Run(test.name, checkRejected)
	}
}

// TestReusesExtractedSection verifies every link to the same section, from
// several documents and link forms, shares one published file.
func (sectionsSuite) TestReusesExtractedSection(c *tc.C) {
	source := sectionSource()
	source["HARNESS.md"].Data = []byte("# Harness\n\n" +
		"[Errors](go/guide.md#errors) and [again][e].\n\n" +
		"[e]: go/guide.md#errors\n")
	source["go/details.md"].Data = []byte(
		"# Details\n\nSee [errors](guide.md#errors).\n",
	)
	bundle, err := Build(source)
	c.Assert(err, tc.IsNil)
	c.Check(publishedSections(bundle), tc.DeepEquals, []string{
		"sections/go/guide/errors.md",
	})
	c.Check(
		string(bundle.Files["HARNESS.md"]),
		tc.Equals,
		"# Harness\n\n"+
			"[Errors](sections/go/guide/errors.md) and [again][e].\n\n"+
			"[e]: sections/go/guide/errors.md\n",
	)
	c.Check(
		string(bundle.Files["go/details.md"]),
		tc.Equals,
		"# Details\n\nSee [errors](../sections/go/guide/errors.md).\n",
	)
}

// TestRewritesLinkForms verifies each supported destination form is
// rewritten in place and code examples are left unchanged.
func (sectionsSuite) TestRewritesLinkForms(c *tc.C) {
	cases := []struct {
		markdown string
		name     string
		want     string
	}{
		{
			name:     "inline",
			markdown: "[Errors](go/guide.md#errors)",
			want:     "[Errors](sections/go/guide/errors.md)",
		},
		{
			name:     "angle brackets",
			markdown: "[Errors](<go/guide.md#errors>)",
			want:     "[Errors](<sections/go/guide/errors.md>)",
		},
		{
			name:     "title",
			markdown: "[Errors](go/guide.md#errors \"Errors\")",
			want:     "[Errors](sections/go/guide/errors.md \"Errors\")",
		},
		{
			name:     "escaped",
			markdown: "[Errors](go/guide\\(1\\).md#errors)",
			want:     "[Errors](sections/go/guide%281%29/errors.md)",
		},
		{
			name:     "query",
			markdown: "[Errors](go/guide.md?view=raw#errors)",
			want:     "[Errors](sections/go/guide/errors.md?view=raw)",
		},
		{
			name:     "image label",
			markdown: "[![Asset](assets/raw.bin)](go/guide.md#errors)",
			want: "[![Asset](assets/raw.bin)]" +
				"(sections/go/guide/errors.md)",
		},
		{
			name:     "reference",
			markdown: "[Errors][e]\n\n[e]: go/guide.md#errors",
			want:     "[Errors][e]\n\n[e]: sections/go/guide/errors.md",
		},
		{
			name: "code examples",
			markdown: "`[x](go/guide.md#errors)` " +
				"[Errors](go/guide.md#errors)\n\n" +
				"```md\n[x](go/guide.md#errors)\n```",
			want: "`[x](go/guide.md#errors)` " +
				"[Errors](sections/go/guide/errors.md)\n\n" +
				"```md\n[x](go/guide.md#errors)\n```",
		},
	}
	for _, test := range cases {
		checkRewritten := func(c *testing.T) {
			source := sectionSource()
			source["HARNESS.md"].Data = []byte(
				"# Harness\n\n" + test.markdown + "\n",
			)
			source["go/guide(1).md"] = &fstest.MapFile{
				Data: source["go/guide.md"].Data,
			}
			bundle, err := Build(source)
			tc.Assert(c, err, tc.IsNil)
			tc.Check(
				c,
				string(bundle.Files["HARNESS.md"]),
				tc.Equals,
				"# Harness\n\n"+test.want+"\n",
			)
		}
		c.Run(test.name, checkRewritten)
	}
}
