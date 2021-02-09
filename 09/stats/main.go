package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"reflect"
	"time"
)

func main() {
	initFiles()
	stats()
	change()
}

const filename = "test.txt"

func initFiles() {
	file, err := os.Create(filename)
	exitIfError(err)
	defer file.Close()
	io.WriteString(file, "New file content\n")
}

func stats() {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		log.Fatalf("Not found file: %s\n", filename)
	}
	exitIfError(err)
	fmt.Printf("%T = %+v\n\n", info, info)

	internal := info.Sys()
	fmt.Printf("%T = %+v\n", internal, internal)
}

func change() {
	err := os.Chmod(filename, 0644)
	exitIfError(err)
	err = os.Chown(filename, os.Getuid(), os.Getgid())
	exitIfError(err)
	err = os.Chtimes(filename, time.Now(), time.Now())
	exitIfError(err)
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%s: %+v\n", reflect.TypeOf(err), errors.WithStack(err))
	}
}
