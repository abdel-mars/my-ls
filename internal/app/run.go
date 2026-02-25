package app

import (
	"fmt"
	// "log"
	"os"
	"path/filepath"
	"my-ls/internal/flags"
)


func ProcessPath(path string) {

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("my-ls" , err) 
		return;
	}
	fmt.Println(path + ":" )
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
	for _, entry := range entries {
		if entry.IsDir() {
			ProcessPath(filepath.Join(path, entry.Name()))
		}
	}
	

}

func ListOneLevel(path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("my-ls", err)
		return
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}

func Run(args []string) {
	parsedFlags, paths := flags.Parse(args)

	if len(paths) == 0 {
		paths = []string{"."}
	}
	for _, path := range paths {
		if parsedFlags.Recursive {
			ProcessPath(path)
		} else {

		}
	}

}


// Parse flags
// If no paths → default to "."
// If Recursive → use ProcessPath
// Else → only list one level (no recursion)

