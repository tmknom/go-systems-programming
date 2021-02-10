package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Page size: %d\n", os.Getpagesize())
}
