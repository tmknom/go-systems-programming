package main

import (
	"fmt"
	"sync"
)

func main() {
	smap := sync.Map{}

	// なんでも入る
	smap.Store("hello", "world")
	smap.Store(1, 2)

	// 取り出す
	value, ok := smap.Load("hello")
	fmt.Printf("key=%v, value=%v, exists?=%v\n", "hello", value, ok)

	// 標準のfor文は使えない
	smap.Range(func(key, value interface{}) bool {
		fmt.Printf("key=%v, value=%v\n", key, value)
		return true
	})

	// 未登録の場合は登録して、その値を取得
	actual, loaded := smap.LoadOrStore(1, 3)
	fmt.Printf("actual=%v, loaded?=%v\n", actual, loaded)
	actual, loaded = smap.LoadOrStore(2, 4)
	fmt.Printf("actual=%v, loaded?=%v\n", actual, loaded)
}
