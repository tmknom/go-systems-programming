package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	exercises2()
}

func exercises2() {
	const filename = "test.csv"
	file, _ := os.Create(filename)
	writer := csv.NewWriter(file)
	writer.Write([]string{"\"foo", "bar", "baz"})
	writer.Write([]string{"1", "2", "3", "4"})
	writer.Flush()
	file.Close()
	os.Remove(filename)
}

func exercises1() {
	fmt.Fprintf(os.Stdout, "format: %s\n", "文字列だよ")
	fmt.Fprintf(os.Stdout, "format: %d\n", 100)
	fmt.Fprintf(os.Stdout, "format: %f\n", 3.14)
}

func text() {
	const filename = "test.txt"
	const multiFilename = "multi.txt"
	const gzipFilename = "test.txt.gz"

	// ファイル
	file, _ := os.Create(filename)
	file.Write([]byte("os.File example\n"))
	file.Close()

	// 標準出力
	os.Stdout.Write([]byte("os.Stdout example\n"))

	// バッファ
	var buf bytes.Buffer
	buf.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buf.String())

	// TCP
	conn, err := net.Dial("tcp", "127.0.0.1:8082")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.1\nHost: 127.0.0.1:8082\n\n")
	io.Copy(os.Stdout, conn)

	// MultiWriter
	multiFile, _ := os.Create(multiFilename)
	multiWriter := io.MultiWriter(multiFile, os.Stdout)
	io.WriteString(multiWriter, "io.MultiWriter example\n")
	multiFile.Close()

	// gzip
	gzipFile, _ := os.Create(gzipFilename)
	gzipWriter := gzip.NewWriter(gzipFile)
	gzipWriter.Header.Name = "test.txt"
	io.WriteString(gzipWriter, "gzip.Writer example\n")
	gzipWriter.Close()

	// bufio
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()

	// format
	fmt.Fprintf(os.Stdout, "format: %v\n", time.Now())

	// json
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})

	// ゴミ掃除
	os.Remove(filename)
	os.Remove(multiFilename)
	os.Remove(gzipFilename)
}
