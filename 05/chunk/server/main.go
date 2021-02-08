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
	"strings"
)

func processSession(conn net.Conn) {
	for {
		// コネクションが確立されてリクエストが飛んできたらデータを読み取る
		request, err := http.ReadRequest(bufio.NewReader(conn))

		if err != nil {
			if err == io.EOF {
				// ソケットがクローズされた場合は何もせず正常終了
				log.Println("socket closed, Goodbye!")
				break
			} else {
				exitIfError(err)
			}
		}

		// 受け取ったデータを標準出力
		dump, err := httputil.DumpRequest(request, true)
		exitIfError(err)
		fmt.Printf("%s\n\n", string(dump))

		// レスポンスを書き込む
		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))

		for i := 0; i < 20; i++ {
			content := fmt.Sprintf("foo_%d\n", i*i)
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
	}
}

func main() {
	// ソケットのリッスン開始
	listener, err := net.Listen("tcp", "localhost:8888")
	exitIfError(err)
	log.Printf("Server is running at %s\n", listener.Addr().String())

	for {
		// クライアントからのコネクションを待つ
		conn, err := listener.Accept()
		exitIfError(err)
		defer conn.Close()
		log.Printf("Accept %v\n", conn.RemoteAddr())
		go processSession(conn)
	}
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
