package publication

import (
	"fmt"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
	"slices"
)

// Write creates destination and writes the bundle without overwriting an
// existing directory, file, or symlink. Its parent must already exist.
// Invalid file paths are rejected before destination is created. Other
// failures may leave a partial new directory; the caller owns its cleanup.
//
// The following errors may be expected:
//   - [fs.ErrExist] the destination already exists.
//   - [fs.ErrNotExist] the destination's parent does not exist.
func Write(bundle Bundle, destination string) error {
	for name := range bundle.Files {
		if !fs.ValidPath(name) || name == "." {
			return fmt.Errorf("invalid output path %q", name)
		}
		_, err := filepath.Localize(name)
		if err != nil {
			return fmt.Errorf(
				"localising output path %q: %w",
				name,
				err,
			)
		}
	}
	err := os.Mkdir(destination, 0o755)
	if err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}
	root, err := os.OpenRoot(destination)
	if err != nil {
		return fmt.Errorf("opening output directory: %w", err)
	}
	defer root.Close()
	for _, name := range slices.Sorted(maps.Keys(bundle.Files)) {
		local, err := filepath.Localize(name)
		if err != nil {
			return fmt.Errorf("localising %q: %w", name, err)
		}
		err = root.MkdirAll(filepath.Dir(local), 0o755)
		if err != nil {
			return fmt.Errorf(
				"creating parent of %q: %w",
				name,
				err,
			)
		}
		err = root.WriteFile(local, bundle.Files[name], 0o644)
		if err != nil {
			return fmt.Errorf("writing %q: %w", name, err)
		}
	}
	return nil
}
