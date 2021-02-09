package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
)

func main() {
	open()
	read()
	append()
	read()
}

const filename = "test.txt"

func open() {
	file, err := os.Create(filename)
	exitIfError(err)
	defer file.Close()
	io.WriteString(file, "New file content\n")
}

func read() {
	file, err := os.Open(filename)
	exitIfError(err)
	defer file.Close()
	fmt.Printf("Read file:\n")
	io.Copy(os.Stdout, file)
}

func append() {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)
	exitIfError(err)
	defer file.Close()
	io.WriteString(file, "Appended content\n")
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
