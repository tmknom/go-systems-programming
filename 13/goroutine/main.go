package main

import (
	"fmt"
	"time"
)

func sub(c int) {
	fmt.Printf("share by arguments: %d\n", c*c)
}

func main() {
	variable()
	closure()
	loop()
	time.Sleep(time.Second)
}

func variable() {
	go sub(10)
}

func closure() {
	c := 20
	go func() {
		fmt.Printf("share by arguments: %d\n", c*c)
	}()
}

func loop() {
	tasks := []string{
		"first",
		"second",
		"third",
	}
	for _, task := range tasks {
		// goroutineが起動するときにはループが回りきって
		// すべてのgoroutineで最後の変数を参照してしまう
		go func() {
			fmt.Printf("loop: %s\n", task)
		}()
	}
}
