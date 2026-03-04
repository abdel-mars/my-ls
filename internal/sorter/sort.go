package sorter

import (
	"io/fs"
	flagpkg "my-ls/internal/flags"
)


func BubbleSort(entries []fs.DirEntry, flags flagpkg.Flags) {
	n := len(entries)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {

			swap := false

			if flags.TimeSort {
				infoJ, _ := entries[j].Info()
				infoJ1, _ := entries[j+1].Info()

				if flags.Reverse {
					// -tr → oldest first
					if infoJ.ModTime().After(infoJ1.ModTime()) {
						swap = true
					}
				} else {
					// -t → newest first
					if infoJ.ModTime().Before(infoJ1.ModTime()) {
						swap = true
					}
				}

			} else {
				// Alphabetical
				if flags.Reverse {
					if entries[j].Name() < entries[j+1].Name() {
						swap = true
					}
				} else {
					if entries[j].Name() > entries[j+1].Name() {
						swap = true
					}
				}
			}

			if swap {
				entries[j], entries[j+1] = entries[j+1], entries[j]
			}
		}
	}
}

func SortEntries(entries []fs.DirEntry, flags flagpkg.Flags) []fs.DirEntry {
	BubbleSort(entries, flags)
	return entries
}

