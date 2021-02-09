package main

import (
	"fmt"
	"github.com/edsrzf/mmap-go"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// テストデータをファイル書き込み
	testData := []byte("0123456789ABCDEF")
	testPath := filepath.Join(os.TempDir(), "testdata")
	err := ioutil.WriteFile(testPath, testData, 0644)
	exitIfError(err)

	// メモリにマッピング
	// mは[]byteのエイリアスなので添字アクセス可能
	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
	exitIfError(err)
	defer f.Close()
	m, err := mmap.Map(f, mmap.COPY, 0)
	exitIfError(err)
	defer m.Unmap()

	// メモリ上のデータを修正して書き込む
	m[9] = 'X'
	m.Flush()

	// 読み込んでみる
	fileData, err := ioutil.ReadAll(f)
	exitIfError(err)
	fmt.Printf("original: %s\n", testData)
	fmt.Printf("mmap: %s\n", m)
	fmt.Printf("file: %s\n", fileData)
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
