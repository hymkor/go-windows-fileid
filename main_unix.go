// +build !windows

package fileid

import (
	"errors"
	"os"
	"syscall"
)

func queryFilenameById(fname string) (uint64, error) {
	fi, err := os.Stat(fname)
	if err != nil {
		return 0, err
	}
	stat, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return 0, errors.New("os.Fileinfo.Sys() is not syscall.Stat_t")
	}
	return stat.Ino, nil
}
