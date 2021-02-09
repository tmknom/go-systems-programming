package main

import (
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
)

func main() {
	initFiles()
	rename()
	//remove()
}

const filename = "test.txt"
const dirname = "test_dir"

func initFiles() {
	file, err := os.Create(filename)
	exitIfError(err)
	defer file.Close()
	io.WriteString(file, "New file content\n")
	os.MkdirAll(dirname+"/foo/bar/baz", 0755)
}

func remove() {
	os.Remove(filename)
	os.RemoveAll(dirname)
}

func truncate() {
	err := os.Truncate(filename, 8)
	exitIfError(err)
}

func rename() {
	renamed := "renamed_" + filename
	err := os.Rename(filename, renamed)
	exitIfError(err)
	err = os.Rename(renamed, dirname+"/"+renamed)
	exitIfError(err)
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
