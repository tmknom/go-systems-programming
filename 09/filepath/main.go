package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	join()
	split()
	fragments()
	list()
	clean()
	env()
	tilda()
	fmt.Printf("~/code: %s\n", cleanPath("~/code"))
	fmt.Printf("${GOPATH}/code: %s\n", cleanPath("${GOPATH}/code"))

}

func cleanPath(path string) string {
	if len(path) > 1 && path[0:2] == "~/" {
		my, err := user.Current()
		exitIfError(err)
		path = my.HomeDir + path[1:]
	}
	path = os.ExpandEnv(path)
	return filepath.Clean(path)
}

func tilda() {
	my, err := user.Current()
	exitIfError(err)
	fmt.Printf("home dir: %s\n", my.HomeDir)
}

func env() {
	fmt.Printf("env: %s\n", os.ExpandEnv("${GOPATH}/src/github.com/"))
}

func clean() {
	// パスをそのままクリーンにする
	fmt.Printf("clean: %s\n", filepath.Clean("./path/foo/../path.go"))

	abs, err := filepath.Abs("./path/foo/../path.go")
	exitIfError(err)
	fmt.Printf("abs: %s\n", abs)

	rel, err := filepath.Rel("/usr/local/go/src", "/usr/local/go/src/path/foo/path.go")
	exitIfError(err)
	fmt.Printf("rel: %s\n", rel)
}

func list() {
	for _, path := range filepath.SplitList(os.Getenv("PATH")) {
		fmt.Printf("%+v\n", path)
	}
}

func fragments() {
	split := strings.Split(os.Getenv("GOPATH"), string(filepath.Separator))
	fmt.Printf("%+v\n", split)
}

func split() {
	dir, name := filepath.Split(os.Getenv("GOPATH"))
	fmt.Printf("Dir: %s, Name: %s\n", dir, name)
}

func join() {
	fmt.Printf("file path: %s\n", filepath.Join(os.TempDir(), "test.txt"))
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
