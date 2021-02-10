package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	err := cmd.Run()
	exitIfError(err)

	state := cmd.ProcessState
	fmt.Printf("state %s\n", state.String())
	fmt.Printf("  pid: %d\n", state.Pid())
	fmt.Printf("  exited: %+v\n", state.Exited())
	fmt.Printf("  success: %d\n", state.Success())
	fmt.Printf("  system: %+v\n", state.SystemTime())
	fmt.Printf("  user: %+v\n", state.UserTime())
	fmt.Printf("  state %+v\n", state)
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
