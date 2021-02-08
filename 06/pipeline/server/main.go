package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// 順番に従ってconnに書き出す
// goroutineでの実行を想定
func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	log.Printf("writeToConn start\n")
	defer conn.Close()
	for sessionResponse := range sessionResponses {
		log.Printf("writeToConn wait sessionResponse...\n")
		response := <-sessionResponse
		response.Write(conn)
		close(sessionResponse)
		response.Body.Close()
		log.Printf("writeToConn done write\n")
	}
}

func handleRequest(request *http.Request, resultReceiver chan *http.Response) {
	log.Printf("handleRequest start: %s\n", request.RequestURI)

	// 受け取ったデータを標準出力
	_, err := httputil.DumpRequest(request, true)
	exitIfError(err)
	//fmt.Printf(string(dump))

	// レスポンスを書き込む
	// セッションを維持するため、Keep-Aliveでないといけない
	content := fmt.Sprintf("Hello, World %s", request.RequestURI)
	response := &http.Response{
		StatusCode:    200,
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: int64(len(content)),
		Body:          ioutil.NopCloser(strings.NewReader(content)),
	}

	// 処理が終わったらチャネルに書き込み
	// ブロックされていたwriteToConnの処理を再始動する
	resultReceiver <- response

	log.Printf("handleRequest enqueue response(%s) to resultReceiver\n", content)
}

func processSession(conn net.Conn) {
	log.Printf("processSession start\n")

	// セッション内のリクエストを順に処理するためのチャネル
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)

	// レスポンスを直列化してソケットに書き出す
	go writeToConn(sessionResponses, conn)

	reader := bufio.NewReader(conn)
	for {
		// タイムアウトの設定
		err := conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		exitIfError(err)

		// コネクションが確立されてリクエストが飛んできたらデータを読み取る
		request, err := http.ReadRequest(reader)
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				// タイムアウトの場合はログを吐いて正常終了
				log.Println("Oops...🐶 Timeout!")
				break
			} else if err == io.EOF {
				// ソケットがクローズされた場合は何もせず正常終了
				log.Println("socket closed, Goodbye!")
				break
			} else {
				exitIfError(err)
			}
		}

		sessionResponse := make(chan *http.Response, 50)
		sessionResponses <- sessionResponse

		// 非同期でレスポンスを実行
		go handleRequest(request, sessionResponse)
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
