package app

import (
	"fmt"
	"strings"
	flagpkg "my-ls/internal/flags"
	"os"
	"path/filepath"
	"my-ls/internal/sorter"
)


func ProcessPath(path string, flags flagpkg.Flags) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("my-ls", err)
		return
	}

	filtered := []os.DirEntry{}
	for _, e := range entries {
		if flags.All || !strings.HasPrefix(e.Name(), ".") {
			filtered = append(filtered, e)
		}
	}

	filtered = sorter.SortEntries(filtered, flags)

	fmt.Println(path + ":")
	for _, entry := range filtered {
		fmt.Println(entry.Name())
	}

	// Recurse into subdirectories
	for _, entry := range filtered {
		if entry.IsDir() {
			ProcessPath(filepath.Join(path, entry.Name()), flags)
		}
	}
}


func ListOneLevel(path string, flags flagpkg.Flags) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("my-ls", err)
		return
	}

	filtered := []os.DirEntry{}
	for _, e := range entries {
		if flags.All || !strings.HasPrefix(e.Name(), ".") {
			filtered = append(filtered, e)
		}
	}

	filtered = sorter.SortEntries(filtered, flags)

	for _, entry := range filtered {
		fmt.Println(entry.Name())
	}
}

func Run(args []string) {
	parsedFlags, paths := flagpkg.Parse(args)

	if len(paths) == 0 {
		paths = []string{"."}
	}
	for _, path := range paths {
		if parsedFlags.Recursive {
			ProcessPath(path, parsedFlags)
		} else {
			// fmt.Println("marsWasHere")
			ListOneLevel(path, parsedFlags)
		}
	}

}
