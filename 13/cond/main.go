package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	for _, name := range []string{"A", "B", "C"} {
		go func(name string) {
			// ロックしてからWaitメソッドを呼ぶ
			mutex.Lock()
			defer mutex.Unlock()
			// Broadcast()が呼ばれるまで待つ
			fmt.Printf("wait %s\n", name)
			cond.Wait()
			// 呼ばれた！
			fmt.Printf("called %s\n", name)
		}(name)
	}
	fmt.Printf("start!\n")
	time.Sleep(time.Second)
	cond.Broadcast()
	time.Sleep(time.Second)
}
