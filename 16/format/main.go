package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
	fmt.Printf("%s\n", now.Format("2006/01/02 15:04:05 MST"))
}
