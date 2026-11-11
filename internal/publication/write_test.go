package publication_test

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/juju/tc"
	"github.com/tlm/thebook/internal/publication"
)

// writeSuite exercises output creation, permissions and destination safety.
type writeSuite struct{}

// TestWrite registers the public Write contract tests.
func TestWrite(t *testing.T) {
	tc.Run(t, &writeSuite{})
}

// TestCreatesBundle verifies complete output, exact bytes and permissions.
func (writeSuite) TestCreatesBundle(c *tc.C) {
	bundle, err := publication.Build(sourceFixture())
	c.Assert(err, tc.IsNil)
	destination := filepath.Join(c.TempDir(), "published")
	err = publication.Write(bundle, destination)
	c.Assert(err, tc.IsNil)
	for name, want := range bundle.Files {
		path := filepath.Join(destination, filepath.FromSlash(name))
		got, err := os.ReadFile(path)
		c.Assert(err, tc.IsNil)
		c.Check(bytes.Equal(got, want), tc.Equals, true)
		info, err := os.Stat(path)
		c.Assert(err, tc.IsNil)
		c.Check(info.Mode().Perm(), tc.Equals, fs.FileMode(0644))
	}
	for _, name := range []string{".", "assets", "go"} {
		info, err := os.Stat(filepath.Join(destination, name))
		c.Assert(err, tc.IsNil)
		c.Check(info.IsDir(), tc.Equals, true)
		c.Check(info.Mode().Perm(), tc.Equals, fs.FileMode(0755))
	}
}

// TestInvalidPaths verifies paths are checked before creating output.
func (writeSuite) TestInvalidPaths(c *tc.C) {
	for _, name := range []string{
		"", ".", "..", "../escape.txt", "/absolute.txt",
		"nested/../../escape.txt", "nested/../escape.txt",
		"nested/./file.txt", "nested//file.txt", "nested/",
	} {
		checkInvalidPath := func(c *testing.T) {
			parent := c.TempDir()
			destination := filepath.Join(parent, "published")
			bundle := publication.Bundle{
				Files: map[string][]byte{
					"HARNESS.md": []byte("# Harness\n"),
					name:         []byte("blocked"),
				},
				Version: "invalid-path-fixture",
			}
			err := publication.Write(bundle, destination)
			tc.Assert(c, err == nil, tc.Equals, false)
			_, err = os.Lstat(destination)
			tc.Check(
				c,
				errors.Is(err, fs.ErrNotExist),
				tc.Equals,
				true,
			)
			entries, err := os.ReadDir(parent)
			tc.Assert(c, err, tc.IsNil)
			tc.Check(c, len(entries), tc.Equals, 0)
		}
		c.Run(name, checkInvalidPath)
	}
}

// TestMissingParent verifies Write does not create a missing parent tree.
func (writeSuite) TestMissingParent(c *tc.C) {
	bundle, err := publication.Build(sourceFixture())
	c.Assert(err, tc.IsNil)
	parent := filepath.Join(c.TempDir(), "missing")
	err = publication.Write(bundle, filepath.Join(parent, "published"))
	c.Assert(err == nil, tc.Equals, false)
	_, err = os.Lstat(parent)
	c.Check(errors.Is(err, fs.ErrNotExist), tc.Equals, true)
}

// TestRefusesExistingDirectory verifies existing output remains unchanged.
func (writeSuite) TestRefusesExistingDirectory(c *tc.C) {
	bundle, err := publication.Build(sourceFixture())
	c.Assert(err, tc.IsNil)
	destination := filepath.Join(c.TempDir(), "published")
	err = os.Mkdir(destination, 0755)
	c.Assert(err, tc.IsNil)
	path := filepath.Join(destination, "HARNESS.md")
	err = os.WriteFile(path, []byte("original"), 0644)
	c.Assert(err, tc.IsNil)
	err = publication.Write(bundle, destination)
	c.Assert(err == nil, tc.Equals, false)
	entries, err := os.ReadDir(destination)
	c.Assert(err, tc.IsNil)
	c.Check(len(entries), tc.Equals, 1)
	data, err := os.ReadFile(path)
	c.Assert(err, tc.IsNil)
	c.Check(string(data), tc.Equals, "original")
}

// TestRefusesExistingEmptyDirectory verifies even empty output is refused.
func (writeSuite) TestRefusesExistingEmptyDirectory(c *tc.C) {
	bundle, err := publication.Build(sourceFixture())
	c.Assert(err, tc.IsNil)
	destination := filepath.Join(c.TempDir(), "published")
	err = os.Mkdir(destination, 0755)
	c.Assert(err, tc.IsNil)
	err = publication.Write(bundle, destination)
	c.Assert(err == nil, tc.Equals, false)
	entries, err := os.ReadDir(destination)
	c.Assert(err, tc.IsNil)
	c.Check(len(entries), tc.Equals, 0)
}

// TestRefusesExistingFile verifies an existing file is not replaced.
func (writeSuite) TestRefusesExistingFile(c *tc.C) {
	bundle, err := publication.Build(sourceFixture())
	c.Assert(err, tc.IsNil)
	destination := filepath.Join(c.TempDir(), "published")
	err = os.WriteFile(destination, []byte("original"), 0644)
	c.Assert(err, tc.IsNil)
	err = publication.Write(bundle, destination)
	c.Assert(err == nil, tc.Equals, false)
	data, err := os.ReadFile(destination)
	c.Assert(err, tc.IsNil)
	c.Check(string(data), tc.Equals, "original")
}

// TestRefusesSymlinkDestination verifies existing links cannot redirect writes.
func (writeSuite) TestRefusesSymlinkDestination(c *tc.C) {
	bundle, err := publication.Build(sourceFixture())
	c.Assert(err, tc.IsNil)
	parent := c.TempDir()
	target := filepath.Join(parent, "target")
	err = os.Mkdir(target, 0755)
	c.Assert(err, tc.IsNil)
	destination := filepath.Join(parent, "published")
	err = os.Symlink(target, destination)
	c.Assert(err, tc.IsNil)
	err = publication.Write(bundle, destination)
	c.Assert(err == nil, tc.Equals, false)
	entries, err := os.ReadDir(target)
	c.Assert(err, tc.IsNil)
	c.Check(len(entries), tc.Equals, 0)
	link, err := os.Readlink(destination)
	c.Assert(err, tc.IsNil)
	c.Check(link, tc.Equals, target)
}
