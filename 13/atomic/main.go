package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var id int64

func generateId() int64 {
	return atomic.AddInt64(&id, 1)
}

func main() {
	var wg sync.WaitGroup

	// 完了を待つgoroutineの数
	wg.Add(100)

	// goroutineの実行
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId())
			wg.Done()
		}()
	}

	// すべてのgoroutine完了を待つ
	wg.Wait()
	fmt.Printf("finished!\n")
}
