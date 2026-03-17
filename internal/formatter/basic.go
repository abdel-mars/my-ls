package formatter

import (
	"fmt"
	"os"
)

func PrintBasic(entries []os.DirEntry) {
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}
