package app

import (
	"fmt"
	// "log"
	"os"
	"path/filepath"
)


func ProcessPath(path string) {

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("my-ls" , err)
	}
	fmt.Println(path + "/" )
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
	for _, entry := range entries {
		if entry.IsDir() {
			ProcessPath(filepath.Join(path, entry.Name()))
		}
	}
	

}


func Run(args []string) {
	if len(args) == 0 {
		ProcessPath(".")
	} else {
		for _, arg := range args {
			ProcessPath(arg)
		}
	}

// 	If no args → use "."
// Else → for each arg → call ProcessPath(arg)
}

