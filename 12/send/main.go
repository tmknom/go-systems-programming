package main

import (
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	pid, err := strconv.Atoi(os.Args[1])
	exitIfError(err)

	process, err := os.FindProcess(pid)
	exitIfError(err)

	err = process.Signal(os.Kill)
	exitIfError(err)

	//err = process.Kill()
	//exitIfError(err)
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
