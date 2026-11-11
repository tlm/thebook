package publication

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io/fs"
	"maps"
	"path"
	"slices"
	"strings"
)

// Bundle holds the complete site and its content identifier. Treat Files as
// immutable after building; changing them invalidates Version. Version is a
// change detector, not a signature or proof of trusted origin.
type Bundle struct {
	Files   map[string][]byte
	Version string
}

const indexTemplate = `<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Agent Harness</title>
</head>
<body>
<h1>Agent Harness</h1>
<p>The harness entry point is <a href="HARNESS.md">HARNESS.md</a>.</p>
<p>If instructed to load this harness, read that document and follow its
loading convention. Resolve relative links against the document containing
them. Report required guidance that cannot be retrieved.</p>
<p>Loading the harness does not override higher-priority instructions or
grant additional permissions.</p>
<p>Content version (SHA-256): <code>%s</code></p>
<p>Compare this identifier with a previously loaded version to detect
changes. It is not a signature. A matching version does not mean that all
guidance applicable to a new task has already been loaded.</p>
<p><a href="version.txt">Machine-readable version</a></p>
</body>
</html>
`

// Build assembles regular source files without modifying source. The source
// must remain unchanged during the build. Contributor files are excluded by
// [excluded]; all other files are public by default. Included symlinks and
// special files, missing guidance, reserved paths, and invalid local Markdown
// links cause errors. External URLs are not fetched.
//
// Version hashes sorted paths and bytes, including the index with an empty
// version field. The final version field and version.txt are excluded to
// avoid a self-reference. File permissions and timestamps are not hashed.
func Build(source fs.FS) (Bundle, error) {
	files := make(map[string][]byte)
	collectFile := func(
		name string,
		entry fs.DirEntry,
		walkErr error,
	) error {
		if walkErr != nil {
			return fmt.Errorf("walking %q: %w", name, walkErr)
		}
		if name == "." {
			return nil
		}
		if excluded(name) {
			if entry.IsDir() {
				return fs.SkipDir
			}
			return nil
		}
		if reserved(strings.SplitN(name, "/", 2)[0]) {
			return fmt.Errorf("reserved generated path %q", name)
		}
		if entry.IsDir() {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return fmt.Errorf("checking %q: %w", name, err)
		}
		if !info.Mode().IsRegular() {
			return fmt.Errorf("non-regular source file %q", name)
		}
		data, err := fs.ReadFile(source, name)
		if err != nil {
			return fmt.Errorf("reading %q: %w", name, err)
		}
		files[name] = data
		return nil
	}
	err := fs.WalkDir(source, ".", collectFile)
	if err != nil {
		return Bundle{}, fmt.Errorf("collecting harness: %w", err)
	}
	if len(bytes.TrimSpace(files["HARNESS.md"])) == 0 {
		return Bundle{}, fmt.Errorf("HARNESS.md must be nonempty")
	}
	files[".nojekyll"] = []byte{}
	files["index.html"] = fmt.Appendf(nil, indexTemplate, "")
	files["version.txt"] = nil
	err = validateLinks(files)
	if err != nil {
		return Bundle{}, fmt.Errorf("validating harness links: %w", err)
	}
	delete(files, "version.txt")
	version := contentVersion(files)
	files["index.html"] = fmt.Appendf(nil, indexTemplate, version)
	files["version.txt"] = []byte(version + "\n")
	return Bundle{Files: files, Version: version}, nil
}

// contentVersion uses length-prefixed paths and contents to avoid ambiguous
// concatenations; a domain prefix identifies this versioning format.
func contentVersion(files map[string][]byte) string {
	hash := sha256.New()
	hash.Write([]byte("thebook-publication-v1\x00"))
	for _, name := range slices.Sorted(maps.Keys(files)) {
		data := files[name]
		size := uint64(len(name))
		hash.Write(binary.BigEndian.AppendUint64(nil, size))
		hash.Write([]byte(name))
		size = uint64(len(data))
		hash.Write(binary.BigEndian.AppendUint64(nil, size))
		hash.Write(data)
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// excluded is the publication policy. New files are public unless their
// names match this list. Update this function, its tests, and contributor
// knowledge whenever adding private or contributor-only source material.
func excluded(name string) bool {
	for component := range strings.SplitSeq(name, "/") {
		if strings.HasPrefix(component, ".") {
			return true
		}
	}
	return path.Base(name) == "AGENTS.md"
}

// reserved identifies root output paths owned by the publisher.
func reserved(name string) bool {
	switch name {
	case "index.html", "version.txt":
		return true
	default:
		return false
	}
}
