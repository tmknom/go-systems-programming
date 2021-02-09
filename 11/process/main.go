package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	pid()
	group()
	user()
	wd()
}

func wd() {
	wd, _ := os.Getwd()
	fmt.Printf("work dir: %s\n", wd)
}

func user() {
	fmt.Printf("user id: %d\n", os.Getuid())
	fmt.Printf("group id: %d\n", os.Getgid())

	groups, _ := os.Getgroups()
	fmt.Printf("sub group ids: %+v\n", groups)

	fmt.Printf("euser id: %d\n", os.Geteuid())
	fmt.Printf("egroup id: %d\n", os.Getegid())
}

func group() {
	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Printf("group id: %d\n", syscall.Getpgrp())
	fmt.Printf("session id: %d\n", sid)
}

func pid() {
	fmt.Printf("process id: %d\n", os.Getpid())
	fmt.Printf("parent process id: %d\n", os.Getppid())
}
