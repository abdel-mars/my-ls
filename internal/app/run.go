package app

import (
	"fmt"
	"my-ls/internal/filesystem"
	flagpkg "my-ls/internal/flags"
	"my-ls/internal/formatter"
	"my-ls/internal/sorter"
	// "os"
	"path/filepath"
)

func ProcessPath(path string, flags flagpkg.Flags) {
	entries, err := filesystem.ReadDirFiltered(path, flags)
	if err != nil {
		fmt.Println("my-ls", err)
		return
	}
	filtered := sorter.SortEntries(entries, flags)

	fmt.Println(path + ":")
	if flags.Long {
		formatter.PrintLong(filtered, path)
	} else {
		formatter.PrintBasic(filtered)
	}

	// Recurse into subdirectories
	for _, entry := range filtered {
		if entry.IsDir() {
			ProcessPath(filepath.Join(path, entry.Name()), flags)
		}
	}
}

func ListOneLevel(path string, flags flagpkg.Flags) {
	entries, err := filesystem.ReadDirFiltered(path, flags)
	if err != nil {
		fmt.Println("my-ls", err)
		return
	}
	filtered := sorter.SortEntries(entries, flags)

	if flags.Long {
		formatter.PrintLong(filtered, path)
	} else {
		formatter.PrintBasic(filtered)
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
