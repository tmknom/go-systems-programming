package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	// 複数回呼んでも、一回しか実行されない
	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}

func initialize() {
	fmt.Printf("initialized!\n")
}
