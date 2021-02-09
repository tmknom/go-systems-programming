package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	traverse()
}

func traverse() {
	base := "../"
	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if info.Name() == "foo" {
				fmt.Printf("skipped %s\n", path)
				return filepath.SkipDir
			}
			return nil
		}
		ext := strings.ToLower(filepath.Ext(info.Name()))
		if ext == ".go" {
			rel, err := filepath.Rel(base, path)
			if err != nil {
				return nil
			}
			fmt.Printf("rel: %s\n", rel)
		}
		return nil
	})
	exitIfError(err)
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
