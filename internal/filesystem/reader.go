package filesystem

import (
	"os"
	"path/filepath"
	"strings"

	flagpkg "my-ls/internal/flags"
)

type dirEntry struct {
	name string
	info os.FileInfo
}

func (d dirEntry) Name() string               { return d.name }
func (d dirEntry) IsDir() bool                { return d.info.IsDir() }
func (d dirEntry) Type() os.FileMode          { return d.info.Mode().Type() }
func (d dirEntry) Info() (os.FileInfo, error) { return d.info, nil }

func ReadDirFiltered(path string, flags flagpkg.Flags) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	filtered := make([]os.DirEntry, 0, len(entries)+2)

	if flags.All {
		if info, err := os.Stat(path); err == nil {
			filtered = append(filtered, dirEntry{name: ".", info: info})
		}
		if info, err := os.Stat(filepath.Join(path, "..")); err == nil {
			filtered = append(filtered, dirEntry{name: "..", info: info})
		}
	}

	for _, e := range entries {
		if flags.All || !strings.HasPrefix(e.Name(), ".") {
			filtered = append(filtered, e)
		}
	}

	return filtered, nil
}
