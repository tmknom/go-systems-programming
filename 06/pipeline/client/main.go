package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {
	// 送信メッセージ
	messages := []string{
		"foo",
		"barbar",
		"bazbazbaz",
	}
	current := 0
	var conn net.Conn = nil
	var err error
	requests := make(chan *http.Request, len(messages))

	// サーバとのコネクションを確立
	conn, err = net.Dial("tcp", "localhost:8888")
	exitIfError(err)
	log.Printf("Access: %d\n", current)
	defer conn.Close()

	// リクエストだけ先に送る
	for i, message := range messages {
		last := i == len(messages)-1

		// リクエストの生成
		request, err := http.NewRequest(
			"GET",
			"http://localhost:8888?message="+message,
			nil,
		)
		exitIfError(err)

		// リクエストヘッダのセット
		if last {
			request.Header.Add("Connection", "close")
		} else {
			request.Header.Add("Connection", "keep-alive")
		}

		// リクエストを送信
		err = request.Write(conn)
		exitIfError(err)
		fmt.Printf("send: %s\n", message)

		requests <- request
	}
	close(requests)

	time.Sleep(1 * time.Second)

	// レスポンスをまとめて受信
	reader := bufio.NewReader(conn)
	for request := range requests {
		// 受信したレスポンスの組み立て
		response, err := http.ReadResponse(reader, request)
		exitIfError(err)

		// レスポンスを標準出力
		dump, err := httputil.DumpResponse(response, true)
		exitIfError(err)
		fmt.Printf("%s\n", string(dump))

		time.Sleep(2 * time.Second)
	}
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
