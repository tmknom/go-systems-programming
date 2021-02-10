package main

import (
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Printf("Accept Ctrl + C for 10 seconds...\n")
	time.Sleep(10 * time.Second)

	// シグナルを無視する
	signal.Ignore(syscall.SIGINT, syscall.SIGTERM)

	fmt.Printf("Ignore Ctrl + C for 10 seconds...\n")
	time.Sleep(10 * time.Second)
}
