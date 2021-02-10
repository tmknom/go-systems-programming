package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "run", "count.go")
	stdout, _ := cmd.StdoutPipe()
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Printf("(stdout) %s\n", scanner.Text())
		}
	}()
	err := cmd.Run()
	exitIfError(err)
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
