package main

import (
	"bytes"
	"errors"
	"flag"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/juju/tc"
)

// commandSuite checks command arguments and publication boundaries.
type commandSuite struct{}

// TestCommand registers command behaviour tests.
func TestCommand(t *testing.T) {
	tc.Run(t, &commandSuite{})
}

// TestBuild writes a complete site and reports its content identifier.
func (commandSuite) TestBuild(c *tc.C) {
	parent := c.TempDir()
	source := filepath.Join(parent, "harness")
	err := os.Mkdir(source, 0o755)
	c.Assert(err, tc.IsNil)
	err = os.WriteFile(
		filepath.Join(source, "HARNESS.md"),
		[]byte("# Harness\n"),
		0o644,
	)
	c.Assert(err, tc.IsNil)
	output := filepath.Join(parent, "site")
	var stdout, stderr bytes.Buffer
	err = run([]string{
		"-source", source, "-output", output,
	}, &stdout, &stderr)
	c.Assert(err, tc.IsNil)
	version, err := os.ReadFile(filepath.Join(output, "version.txt"))
	c.Assert(err, tc.IsNil)
	c.Check(strings.Contains(stdout.String(), string(version)), tc.IsTrue)
	c.Check(stderr.Len(), tc.Equals, 0)
}

// TestHelp reports usage without treating help as a publication failure.
func (commandSuite) TestHelp(c *tc.C) {
	var stdout, stderr bytes.Buffer
	err := run([]string{"-help"}, &stdout, &stderr)
	c.Check(errors.Is(err, flag.ErrHelp), tc.IsTrue)
	c.Check(strings.Contains(stderr.String(), "-source"), tc.IsTrue)
	c.Check(stdout.Len(), tc.Equals, 0)
}

// TestInvalidArguments rejects malformed flags and positional arguments.
func (commandSuite) TestInvalidArguments(c *tc.C) {
	for _, args := range [][]string{{"-unknown"}, {"unexpected"}} {
		checkInvalidArguments := func(c *testing.T) {
			var stdout, stderr bytes.Buffer
			err := run(args, &stdout, &stderr)
			tc.Check(c, err == nil, tc.IsFalse)
			tc.Check(c, stdout.Len(), tc.Equals, 0)
		}
		c.Run(strings.Join(args, " "), checkInvalidArguments)
	}
}

// TestOutputInSource rejects nested output before creating any files.
func (commandSuite) TestOutputInSource(c *tc.C) {
	for _, viaSymlink := range []bool{false, true} {
		checkOutputInSource := func(c *testing.T) {
			parent := c.TempDir()
			source := filepath.Join(parent, "harness")
			err := os.Mkdir(source, 0o755)
			tc.Assert(c, err, tc.IsNil)
			outputParent := source
			if viaSymlink {
				outputParent = filepath.Join(parent, "alias")
				err = os.Symlink(source, outputParent)
				tc.Assert(c, err, tc.IsNil)
			}
			var stdout, stderr bytes.Buffer
			err = run([]string{
				"-source", source,
				"-output", filepath.Join(outputParent, "site"),
			}, &stdout, &stderr)
			tc.Assert(c, err == nil, tc.IsFalse)
			entries, err := os.ReadDir(source)
			tc.Assert(c, err, tc.IsNil)
			tc.Check(c, len(entries), tc.Equals, 0)
		}
		c.Run(strconv.FormatBool(viaSymlink), checkOutputInSource)
	}
}
