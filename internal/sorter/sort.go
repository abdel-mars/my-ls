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
				nameJ := entries[j].Name()
				nameJ1 := entries[j+1].Name()

				rankJ := dotRank(nameJ, flags.All)
				rankJ1 := dotRank(nameJ1, flags.All)

				if rankJ != rankJ1 {
					if flags.Reverse {
						swap = rankJ < rankJ1
					} else {
						swap = rankJ > rankJ1
					}
				} else {
					if flags.Reverse {
						if nameJ < nameJ1 {
							swap = true
						}
					} else {
						if nameJ > nameJ1 {
							swap = true
						}
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

func dotRank(name string, includeDots bool) int {
	if !includeDots {
		return 0
	}
	if name == "." {
		return 0
	}
	if name == ".." {
		return 1
	}
	return 2
}
