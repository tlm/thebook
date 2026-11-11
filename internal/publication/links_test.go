package publication_test

import (
	"testing"
	"testing/fstest"

	"github.com/juju/tc"
	"github.com/tlm/thebook/internal/publication"
)

// linksSuite exercises parsed Markdown destinations and heading anchors.
type linksSuite struct{}

// TestLinks registers Markdown validation contract tests.
func TestLinks(t *testing.T) {
	tc.Run(t, &linksSuite{})
}

// TestAcceptedLinks verifies valid local links and ignored non-local content.
func (linksSuite) TestAcceptedLinks(c *tc.C) {
	cases := []struct {
		markdown string
		name     string
	}{
		{name: "inline", markdown: "[Guide](go/guide.md)"},
		{
			name:     "query",
			markdown: "[Guide](go/guide.md?view=raw#go-guide)",
		},
		{
			name:     "encoded space",
			markdown: "[Notes](go/usage%20notes.md)",
		},
		{
			name: "encoded space and query",
			markdown: "[Notes](go/usage%20notes.md?view=raw" +
				"#usage-notes)",
		},
		{name: "reference", markdown: "[Guide][g]\n\n[g]: go/guide.md"},
		{name: "collapsed reference", markdown: "[Guide][]\n\n" +
			"[Guide]: go/guide.md"},
		{name: "image", markdown: "![Asset](assets/raw.bin)"},
		{name: "reference image", markdown: "![Asset][a]\n\n" +
			"[a]: assets/raw.bin"},
		{name: "same document", markdown: "[Here](#harness)"},
		{name: "heading", markdown: "[Guide](go/guide.md#go-guide)"},
		{
			name:     "formatted heading",
			markdown: "[Guide](go/guide.md#use-go-now)",
		},
		{
			name:     "duplicate heading",
			markdown: "[Guide](go/guide.md#repeat-1)",
		},
		{name: "generated index", markdown: "[Index](index.html)"},
		{name: "generated version", markdown: "[Version](version.txt)"},
		{name: "generated marker", markdown: "[Marker](.nojekyll)"},
		{
			name:     "https",
			markdown: "[Web](https://example.invalid/missing#x)",
		},
		{name: "mail", markdown: "[Mail](mailto:test@example.invalid)"},
		{
			name:     "autolink",
			markdown: "<https://example.invalid/missing>",
		},
		{name: "inline code", markdown: "`[Missing](missing.md)`"},
		{
			name:     "fenced code",
			markdown: "```md\n[Missing](missing.md)\n```",
		},
		{
			name:     "indented code",
			markdown: "    [Missing](missing.md)\n",
		},
		{name: "unused reference", markdown: "[unused]: missing.md\n"},
	}
	for _, test := range cases {
		checkAcceptedLink := func(c *testing.T) {
			source := sourceFixture()
			source["HARNESS.md"].Data = []byte(
				"# Harness\n\n" + test.markdown + "\n",
			)
			source["go/guide.md"].Data = []byte(
				"# Go Guide\n\n## Use *Go* now!\n\n" +
					"## Repeat\n\n## Repeat\n",
			)
			source["go/usage notes.md"] = &fstest.MapFile{
				Data: []byte("# Usage Notes\n"),
			}
			_, err := publication.Build(source)
			tc.Check(c, err, tc.IsNil)
		}
		c.Run(test.name, checkAcceptedLink)
	}
}

// TestEscapedDestinations follows Markdown escape and entity semantics.
func (linksSuite) TestEscapedDestinations(c *tc.C) {
	cases := []struct {
		destination string
		filename    string
	}{
		{destination: `guide\(draft\).md`, filename: "guide(draft).md"},
		{destination: "a&amp;b.md", filename: "a&b.md"},
	}
	for _, test := range cases {
		checkEscapedDestination := func(c *testing.T) {
			source := sourceFixture()
			link := "[Guide](" + test.destination + ")\n"
			source["HARNESS.md"].Data = []byte(
				"# Harness\n\n" + link,
			)
			source[test.filename] = &fstest.MapFile{
				Data: []byte("# Guide\n"),
			}
			_, err := publication.Build(source)
			tc.Check(c, err, tc.IsNil)
		}
		c.Run(test.filename, checkEscapedDestination)
	}
}

// TestLiteralEntityFile cannot satisfy a link that renders differently.
func (linksSuite) TestLiteralEntityFile(c *tc.C) {
	source := sourceFixture()
	source["HARNESS.md"].Data = []byte("# Harness\n\n[Guide](a&amp;b.md)\n")
	source["a&amp;b.md"] = &fstest.MapFile{Data: []byte("# Guide\n")}
	_, err := publication.Build(source)
	c.Check(err == nil, tc.IsFalse)
}

// TestNestedLinks verifies resolution uses the referring document's directory.
func (linksSuite) TestNestedLinks(c *tc.C) {
	source := sourceFixture()
	source["go/guide.md"].Data = []byte(
		"# Go Guide\n\n[Home](../HARNESS.md#harness)\n" +
			"![Asset](../assets/raw.bin)\n" +
			"[Local](./details.md#details)\n",
	)
	source["go/details.md"] = &fstest.MapFile{Data: []byte("# Details\n")}
	_, err := publication.Build(source)
	c.Check(err, tc.IsNil)
}

// TestRejectedLinks verifies missing, excluded, escaping and invalid anchors.
func (linksSuite) TestRejectedLinks(c *tc.C) {
	cases := []struct {
		markdown string
		name     string
	}{
		{name: "missing inline", markdown: "[Missing](missing.md)"},
		{name: "missing reference", markdown: "[Missing][m]\n\n" +
			"[m]: missing.md"},
		{name: "missing image", markdown: "![Missing](missing.png)"},
		{
			name:     "missing reference image",
			markdown: "![Missing][m]\n\n[m]: missing.png",
		},
		{
			name:     "excluded directory",
			markdown: "[Private](.kb/private.md)",
		},
		{name: "excluded file", markdown: "[Agents](AGENTS.md)"},
		{name: "missing local anchor", markdown: "[Here](#absent)"},
		{
			name:     "missing remote anchor",
			markdown: "[Go](go/guide.md#absent)",
		},
		{name: "root escape", markdown: "[Outside](../outside.md)"},
		{
			name:     "encoded root escape",
			markdown: "[Outside](%2e%2e/outside.md)",
		},
		{
			name:     "encoded separator escape",
			markdown: "[Outside](..%2foutside.md)",
		},
		{name: "root relative", markdown: "[Harness](/HARNESS.md)"},
		{
			name:     "normalised escape",
			markdown: "[Outside](go/../../outside.md)",
		},
	}
	for _, test := range cases {
		checkRejectedLink := func(c *testing.T) {
			source := sourceFixture()
			source["HARNESS.md"].Data = []byte(
				"# Harness\n\n" + test.markdown + "\n",
			)
			source[".kb/private.md"] = &fstest.MapFile{
				Data: []byte("# Private\n"),
			}
			source["AGENTS.md"] = &fstest.MapFile{
				Data: []byte("agents"),
			}
			_, err := publication.Build(source)
			tc.Check(c, err == nil, tc.Equals, false)
		}
		c.Run(test.name, checkRejectedLink)
	}
}
