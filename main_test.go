package fileid

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func fsutil(fname string) (uint64, error) {
	bin, err := exec.Command("fsutil", "file", "queryFileId", fname).Output()
	if err != nil {
		return 0, err
	}
	str := string(bin)
	pos := strings.Index(str, "0x")
	if pos < 0 {
		return 0, errors.New("Invalid fsutil output")
	}
	return strconv.ParseUint(strings.Fields(str[pos+2:])[0], 16, 64)
}

func TestQuery(t *testing.T) {
	id, err := Query("main_test.go")
	if err != nil {
		t.Fatalf("QueryFilenameById: %s", err.Error())
	}
	// fmt.Printf("%X\n", id)

	id2, err := fsutil("main_test.go")
	if err != nil {
		t.Fatalf("fsutil file queryFileId: %s", err.Error())
	}

	// fmt.Printf("%X\n", id2)
	if id != id2 {
		t.Fatalf("not matching ID-code: %X != %X", id, id2)
	}
}
