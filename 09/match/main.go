package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"path/filepath"
)

func main() {
	match()
	glob()
}

func glob() {
	files, err := filepath.Glob("./*.go")
	exitIfError(err)
	fmt.Printf("glob: %v\n", files)
}

func match() {
	matched, err := filepath.Match("ma*.go", "main.go")
	exitIfError(err)
	fmt.Printf("match: %v\n", matched)
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
