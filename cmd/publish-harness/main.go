package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/tlm/thebook/internal/publication"
)

func main() {
	err := run(os.Args[1:], os.Stdout, os.Stderr)
	if errors.Is(err, flag.ErrHelp) {
		return
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run assembles the site from a stable source tree. Output must not exist,
// its parent must exist, and it must be outside the source tree. Write
// failures may leave partial output, which is never silently removed.
func run(args []string, stdout io.Writer, stderr io.Writer) error {
	flags := flag.NewFlagSet("publish-harness", flag.ContinueOnError)
	flags.SetOutput(stderr)
	source := flags.String("source", "harness", "harness source directory")
	output := flags.String("output", "_site", "new output directory")
	err := flags.Parse(args)
	if err != nil {
		return fmt.Errorf("parsing arguments: %w", err)
	}
	if flags.NArg() != 0 {
		return fmt.Errorf("unexpected positional arguments")
	}

	sourcePath, err := filepath.EvalSymlinks(*source)
	if err != nil {
		return fmt.Errorf("resolving source: %w", err)
	}
	sourcePath, err = filepath.Abs(sourcePath)
	if err != nil {
		return fmt.Errorf("resolving absolute source: %w", err)
	}
	parent, err := filepath.EvalSymlinks(filepath.Dir(*output))
	if err != nil {
		return fmt.Errorf("resolving output parent: %w", err)
	}
	outputPath, err := filepath.Abs(filepath.Join(
		parent,
		filepath.Base(*output),
	))
	if err != nil {
		return fmt.Errorf("resolving absolute output: %w", err)
	}
	relative, err := filepath.Rel(sourcePath, outputPath)
	if err != nil {
		return fmt.Errorf("checking output location: %w", err)
	}
	if relative == "." ||
		(relative != ".." && !strings.HasPrefix(
			relative,
			".."+string(filepath.Separator),
		)) {
		return fmt.Errorf("output must be outside the source tree")
	}
	root, err := os.OpenRoot(sourcePath)
	if err != nil {
		return fmt.Errorf("opening source: %w", err)
	}
	defer root.Close()
	bundle, err := publication.Build(root.FS())
	if err != nil {
		return fmt.Errorf("building site: %w", err)
	}
	err = publication.Write(bundle, outputPath)
	if err != nil {
		return fmt.Errorf("writing site: %w", err)
	}
	_, err = fmt.Fprintf(
		stdout,
		"Published %d files to %s\nVersion (SHA-256): %s\n",
		len(bundle.Files),
		outputPath,
		bundle.Version,
	)
	if err != nil {
		return fmt.Errorf("reporting publication: %w", err)
	}
	return nil
}
