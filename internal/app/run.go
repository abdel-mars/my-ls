package app

import (
	"fmt"
	// "strings"
	"strings"
	flagpkg "my-ls/internal/flags"
	// "log"
	// "my-ls/internal/flags"
	"os"
	"path/filepath"
)


func ProcessPath(path string, flags flagpkg.Flags) {

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("my-ls" , err) 
		return;
	}
	fmt.Println(path + ":" )
	for _, entry := range entries {
		if !flags.All && strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		fmt.Println(entry.Name())
	}
	for _, entry := range entries {
		if !flags.All && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

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

	for _, entry := range entries {
		if !flags.All && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

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


// Parse flags
// If no paths → default to "."
// If Recursive → use ProcessPath
// Else → only list one level (no recursion)

