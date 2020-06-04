// +build !windows

package fileid

import (
	"errors"
)

func queryFilenameById(fname string) (uint64, error) {
	return 0, errors.New("not support")
}
