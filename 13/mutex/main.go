package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	mutex.Lock()
	id++
	mutex.Unlock()
	return id
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	// 完了を待つgoroutineの数
	wg.Add(100)

	// goroutineの実行
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
			wg.Done()
		}()
	}

	// すべてのgoroutine完了を待つ
	wg.Wait()
	fmt.Printf("finished!\n")
}
