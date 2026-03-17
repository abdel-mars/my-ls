package formatter

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
)

type longRow struct {
	mode  string
	nlink uint64
	user  string
	group string
	size  int64
	mtime string
	name  string
}

func PrintLong(entries []os.DirEntry, basePath string) {
	rows, totalBlocks := buildLongRows(entries, basePath)
	if len(rows) > 0 {
		fmt.Printf("total %d\n", totalBlocks/2)
	}

	widths := computeLongWidths(rows)
	for _, row := range rows {
		fmt.Printf(
			"%s %*d %-*s %-*s %*d %s %s\n",
			row.mode,
			widths.nlink, row.nlink,
			widths.user, row.user,
			widths.group, row.group,
			widths.size, row.size,
			row.mtime,
			row.name,
		)
	}
}

type longWidths struct {
	nlink int
	user  int
	group int
	size  int
}

func computeLongWidths(rows []longRow) longWidths {
	w := longWidths{nlink: 1, user: 1, group: 1, size: 1}
	for _, row := range rows {
		if n := len(strconv.FormatUint(row.nlink, 10)); n > w.nlink {
			w.nlink = n
		}
		if n := len(row.user); n > w.user {
			w.user = n
		}
		if n := len(row.group); n > w.group {
			w.group = n
		}
		if n := len(strconv.FormatInt(row.size, 10)); n > w.size {
			w.size = n
		}
	}
	return w
}

func buildLongRows(entries []os.DirEntry, basePath string) ([]longRow, int64) {
	rows := make([]longRow, 0, len(entries))
	var totalBlocks int64

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Printf("my-ls: %s: %v\n", entry.Name(), err)
			continue
		}

		mode := info.Mode().String()
		size := info.Size()
		modTime := info.ModTime().Local().Format("Jan _2 15:04")

		name := entry.Name()
		if info.Mode()&os.ModeSymlink != 0 {
			linkTarget, err := os.Readlink(filepath.Join(basePath, name))
			if err == nil {
				name = name + " -> " + linkTarget
			}
		}

		nlink := uint64(1)
		uid := uint32(0)
		gid := uint32(0)
		blocks := (size + 511) / 512

		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			nlink = uint64(stat.Nlink)
			uid = stat.Uid
			gid = stat.Gid
			blocks = int64(stat.Blocks)
		}

		totalBlocks += blocks

		userName := lookupUserName(uid)
		groupName := lookupGroupName(gid)

		rows = append(rows, longRow{
			mode:  mode,
			nlink: nlink,
			user:  userName,
			group: groupName,
			size:  size,
			mtime: modTime,
			name:  name,
		})
	}

	return rows, totalBlocks
}

func lookupUserName(uid uint32) string {
	id := strconv.FormatUint(uint64(uid), 10)
	u, err := user.LookupId(id)
	if err != nil || u == nil || u.Username == "" {
		return id
	}
	return u.Username
}

func lookupGroupName(gid uint32) string {
	id := strconv.FormatUint(uint64(gid), 10)
	g, err := user.LookupGroupId(id)
	if err != nil || g == nil || g.Name == "" {
		return id
	}
	return g.Name
}
