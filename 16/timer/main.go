package main

import (
	"fmt"
	"time"
)

func main() {
	after()
	tick()
}

func tick() {
	fmt.Printf("waiting 5 seconds\n")
	for now := range time.Tick(5 * time.Second) {
		fmt.Printf("now: %+v\n", now)
	}
}

func after() {
	fmt.Printf("waiting 5 seconds\n")
	after := time.After(5 * time.Second)
	<-after
	fmt.Printf("done time.After\n")
}
