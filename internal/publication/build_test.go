package publication_test

import (
	"encoding/hex"
	"io/fs"
	"strings"
	"testing"
	"testing/fstest"
	"time"

	"github.com/juju/tc"
	"github.com/tlm/thebook/internal/publication"
)

// buildSuite exercises the publication boundary and generated bundle.
type buildSuite struct{}

// sourceFixture provides linked guidance and a binary asset for publication.
func sourceFixture() fstest.MapFS {
	return fstest.MapFS{
		"HARNESS.md": {
			Data: []byte("# Harness\n\n[Guide](go/guide.md)\n"),
		},
		"go/guide.md": {
			Data: []byte("# Go Guide\n\nUse the harness.\n"),
		},
		"assets/raw.bin": {Data: []byte{0, 255, '\r', '\n', 128}},
	}
}

// TestBuild registers the public Build contract tests.
func TestBuild(t *testing.T) {
	tc.Run(t, &buildSuite{})
}

// TestCopiesAndGenerates verifies exact source bytes and generated artefacts.
func (buildSuite) TestCopiesAndGenerates(c *tc.C) {
	source := sourceFixture()
	bundle, err := publication.Build(source)
	c.Assert(err, tc.IsNil)
	c.Check(len(bundle.Files), tc.Equals, len(source)+3)
	for name, file := range source {
		c.Check(bundle.Files[name], tc.DeepEquals, file.Data)
	}
	decoded, err := hex.DecodeString(bundle.Version)
	c.Assert(err, tc.IsNil)
	c.Check(len(decoded), tc.Equals, 32)
	c.Check(
		string(bundle.Files["version.txt"]),
		tc.Equals,
		bundle.Version+"\n",
	)
	marker, exists := bundle.Files[".nojekyll"]
	c.Check(exists, tc.Equals, true)
	c.Check(len(marker), tc.Equals, 0)
	index := string(bundle.Files["index.html"])
	c.Check(strings.Contains(index, bundle.Version), tc.Equals, true)
	c.Check(strings.Contains(index, `href="HARNESS.md"`), tc.Equals, true)
}

// TestExclusions verifies component exclusions and case-sensitive names.
func (buildSuite) TestExclusions(c *tc.C) {
	source := sourceFixture()
	excluded := []string{
		".kb/private.md", ".git/config", ".github/workflows/pages.yml",
		"go/.kb/private.md", "go/.git/config", "go/.github/config",
		"AGENTS.md", "go/AGENTS.md", ".DS_Store", "go/.DS_Store",
		".ds_store", ".KB/private.txt", "go/.github-notes/file.txt",
		".env", "assets/.private", "assets/.hidden/file.txt",
		".well-known/example.txt",
	}
	for _, name := range excluded {
		source[name] = &fstest.MapFile{
			Data: []byte("[bad](missing.md)"),
		}
	}
	included := []string{
		"agents.md", "go/Agents.md", "notes.kb/file.txt",
	}
	for _, name := range included {
		source[name] = &fstest.MapFile{Data: []byte("public\r\n")}
	}
	bundle, err := publication.Build(source)
	c.Assert(err, tc.IsNil)
	for _, name := range excluded {
		_, exists := bundle.Files[name]
		c.Check(exists, tc.Equals, false)
	}
	for _, name := range included {
		c.Check(bundle.Files[name], tc.DeepEquals, source[name].Data)
	}
}

// TestHiddenMarker verifies source markers are excluded, then generated.
func (buildSuite) TestHiddenMarker(c *tc.C) {
	source := sourceFixture()
	source[".nojekyll"] = &fstest.MapFile{
		Data: []byte("not published"),
	}
	baseline, err := publication.Build(sourceFixture())
	c.Assert(err, tc.IsNil)
	bundle, err := publication.Build(source)
	c.Assert(err, tc.IsNil)
	c.Check(bundle, tc.DeepEquals, baseline)
}

// TestInvalidSources verifies required entry points and reserved root paths.
func (buildSuite) TestInvalidSources(c *tc.C) {
	cases := []struct {
		name   string
		source fstest.MapFS
	}{
		{name: "missing harness", source: fstest.MapFS{}},
		{name: "empty harness", source: fstest.MapFS{
			"HARNESS.md": {Data: []byte{}},
		}},
		{name: "directory harness", source: fstest.MapFS{
			"HARNESS.md": {Mode: fs.ModeDir | 0755},
		}},
	}
	for _, name := range []string{
		"index.html", "version.txt",
	} {
		for _, mode := range []fs.FileMode{0644, fs.ModeDir | 0755} {
			source := sourceFixture()
			source[name] = &fstest.MapFile{Mode: mode}
			cases = append(cases, struct {
				name   string
				source fstest.MapFS
			}{name: name + " " + mode.String(), source: source})
		}
	}
	for _, test := range cases {
		checkInvalidSource := func(c *testing.T) {
			_, err := publication.Build(test.source)
			tc.Check(c, err == nil, tc.Equals, false)
		}
		c.Run(test.name, checkInvalidSource)
	}
}

// TestRejectsNonRegularFiles verifies included special entries are rejected.
func (buildSuite) TestRejectsNonRegularFiles(c *tc.C) {
	for _, mode := range []fs.FileMode{
		fs.ModeSymlink, fs.ModeNamedPipe, fs.ModeSocket, fs.ModeDevice,
	} {
		checkNonRegularFile := func(c *testing.T) {
			source := sourceFixture()
			source["special"] = &fstest.MapFile{
				Data: []byte("HARNESS.md"), Mode: mode | 0644,
			}
			_, err := publication.Build(source)
			tc.Check(c, err == nil, tc.Equals, false)
		}
		c.Run(mode.String(), checkNonRegularFile)
	}
}

// TestSourceUnchangedOnError verifies validation never mutates source bytes.
func (buildSuite) TestSourceUnchangedOnError(c *tc.C) {
	source := sourceFixture()
	source["go/guide.md"].Data = []byte("# Guide\n[Bad](missing.md)\n")
	before := make(fstest.MapFS, len(source))
	for name, file := range source {
		copy := *file
		copy.Data = append([]byte(nil), file.Data...)
		before[name] = &copy
	}
	_, err := publication.Build(source)
	c.Assert(err == nil, tc.Equals, false)
	c.Check(source, tc.DeepEquals, before)
}

// TestVersionChanges verifies included edits, additions, deletions and renames.
func (buildSuite) TestVersionChanges(c *tc.C) {
	baseline, err := publication.Build(sourceFixture())
	c.Assert(err, tc.IsNil)
	cases := []struct {
		mutate func(fstest.MapFS)
		name   string
	}{
		{name: "edit", mutate: func(source fstest.MapFS) {
			source["assets/raw.bin"].Data = []byte{0, 1, 2}
		}},
		{name: "add", mutate: func(source fstest.MapFS) {
			source["extra.txt"] = &fstest.MapFile{
				Data: []byte("extra"),
			}
		}},
		{name: "delete", mutate: func(source fstest.MapFS) {
			delete(source, "assets/raw.bin")
		}},
		{name: "rename", mutate: func(source fstest.MapFS) {
			source["assets/renamed.bin"] = source["assets/raw.bin"]
			delete(source, "assets/raw.bin")
		}},
	}
	for _, test := range cases {
		checkVersionChange := func(c *testing.T) {
			source := sourceFixture()
			test.mutate(source)
			bundle, err := publication.Build(source)
			tc.Assert(c, err, tc.IsNil)
			tc.Check(
				c,
				bundle.Version == baseline.Version,
				tc.Equals,
				false,
			)
		}
		c.Run(test.name, checkVersionChange)
	}
}

// TestVersionStable verifies repeat builds, metadata and excluded edits.
func (buildSuite) TestVersionStable(c *tc.C) {
	source := sourceFixture()
	baseline, err := publication.Build(source)
	c.Assert(err, tc.IsNil)
	repeated, err := publication.Build(source)
	c.Assert(err, tc.IsNil)
	c.Check(repeated, tc.DeepEquals, baseline)
	for _, file := range source {
		file.ModTime = time.Unix(123456789, 0)
	}
	source[".kb/private.md"] = &fstest.MapFile{Data: []byte("private")}
	source["AGENTS.md"] = &fstest.MapFile{Data: []byte("instructions")}
	changed, err := publication.Build(source)
	c.Assert(err, tc.IsNil)
	c.Check(changed, tc.DeepEquals, baseline)
	source[".kb/private.md"].Data = []byte("edited")
	delete(source, "AGENTS.md")
	changed, err = publication.Build(source)
	c.Assert(err, tc.IsNil)
	c.Check(changed, tc.DeepEquals, baseline)

	reordered := make(fstest.MapFS, len(source))
	for _, name := range []string{
		"assets/raw.bin", "go/guide.md", "HARNESS.md", ".kb/private.md",
	} {
		reordered[name] = source[name]
	}
	changed, err = publication.Build(reordered)
	c.Assert(err, tc.IsNil)
	c.Check(changed, tc.DeepEquals, baseline)
}
