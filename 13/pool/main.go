package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	create()
	gc()
}

func create() {
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}
	pool.Put("manualy added: 1")
	pool.Put("manualy added: 2")
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get()) // 新規作成されたモノが取得できる
}

func gc() {
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("gc created: %d", count)
		},
	}
	pool.Put("removed: 1")
	pool.Put("removed: 2")

	// GCを実行するとpoolの中身は失われる
	runtime.GC()

	// 失われる場合とそうでない場合がある
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}
