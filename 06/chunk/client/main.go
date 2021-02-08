package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	// サーバとのコネクションを確立
	conn, err := net.Dial("tcp", "localhost:8888")
	exitIfError(err)
	defer conn.Close()

	// リクエストの生成
	request, err := http.NewRequest(
		"GET",
		"http://localhost:8888",
		nil)

	exitIfError(err)

	// リクエストを送信
	err = request.Write(conn)
	exitIfError(err)

	// レスポンスを標準出力
	reader := bufio.NewReader(conn)
	response, err := http.ReadResponse(reader, request)
	exitIfError(err)

	dump, err := httputil.DumpResponse(response, false)
	exitIfError(err)
	fmt.Printf(string(dump))

	if len(response.TransferEncoding) < 1 || response.TransferEncoding[0] != "chunked" {
		log.Fatalf("wrong transfer encoding: %+v\n", response.TransferEncoding)
	}

	for {
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		// 16進数のサイズをパース
		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if size == 0 {
			break
		}
		exitIfError(err)

		// サイズ数分バッファを確保して読み込み
		line := make([]byte, int(size))
		reader.Read(line)
		reader.Discard(2)
		fmt.Printf("%d bytes: %s", size, string(line))
	}
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
