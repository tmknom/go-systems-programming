package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

func main() {
	// 送信メッセージ
	messages := []string{
		"foo",
		"bar",
		"baz",
	}
	current := 0
	var conn net.Conn = nil

	// リトライ用にループで全体を囲う
	for {
		var err error

		// サーバとのコネクションを確立
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")
			exitIfError(err)
			log.Printf("Access: %d\n", current)
		}

		// リクエストの生成
		request, err := http.NewRequest(
			"POST",
			"http://localhost:8888",
			strings.NewReader(messages[current]),
		)
		exitIfError(err)
		request.Header.Set("Accept-Encoding", "gzip")

		// リクエストを送信
		err = request.Write(conn)
		exitIfError(err)

		// 受信したレスポンスの組み立て
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			log.Printf("retry, caused by '%+v'", err)
			err := conn.Close()
			exitIfError(err)
			conn = nil
			continue
		}
		exitIfError(err)

		// レスポンスを標準出力
		dump, err := httputil.DumpResponse(response, false)
		exitIfError(err)
		fmt.Printf(string(dump))
		defer response.Body.Close()

		if response.Header.Get("Content-Encoding") == "gzip" {
			var buffer bytes.Buffer
			teeReader := io.TeeReader(response.Body, &buffer)

			reader, err := gzip.NewReader(teeReader)
			exitIfError(err)
			io.Copy(os.Stdout, reader)
			reader.Close()

			fmt.Printf("%+v\n", buffer.Bytes())
		} else {
			io.Copy(os.Stdout, response.Body)
		}

		time.Sleep(1 * time.Second)
		current++
		if current == len(messages) {
			break
		}
	}
	err := conn.Close()
	exitIfError(err)
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
