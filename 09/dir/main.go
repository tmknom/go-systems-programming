package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
)

func main() {
	initFiles()
	list()
}

const filename = "test.txt"
const dirname = "test_dir"

func initFiles() {
	err := os.MkdirAll(dirname+"/foo/bar/baz", 0755)
	exitIfError(err)
	file1, err := os.Create(dirname + "/foo_" + filename)
	exitIfError(err)
	defer file1.Close()

	file2, err := os.Create(dirname + "/bar_" + filename)
	exitIfError(err)
	defer file2.Close()

	file3, err := os.Create(dirname + "/baz_" + filename)
	exitIfError(err)
	defer file3.Close()

	io.WriteString(file1, "New file content\n")
	io.WriteString(file2, "New file content\n")
	io.WriteString(file3, "New file content\n")
}

func list() {
	dir, err := os.Open(dirname)
	exitIfError(err)

	fileInfos, err := dir.Readdir(-1) // -1だとすべてのファイルを取得
	exitIfError(err)
	for _, info := range fileInfos {
		if info.IsDir() {
			fmt.Printf("[Dir] %s\n", info.Name())
		} else {
			fmt.Printf("[File] %s\n", info.Name())
		}
	}
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
