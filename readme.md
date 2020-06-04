go-windows-fileid
=================

Query File-ID. File-ID is the number like i-node on Windows.

```go
package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/zetamatta/go-windows-fileid"
)

func mains() error {
	id1, err := fileid.Query("main_test.go")
	if err != nil {
		return err
	}
	fmt.Printf("fileid.Query=%x\n", id1)
	cmd := exec.Command("fsutil", "file", "queryFileId", "main_test.go")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func main() {
	if err := mains(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
```

```
$ go run example.go                                                             fileid.Query=da0000000158ac                                                     ファイル ID は 0x000000000000000000da0000000158ac です      
```
