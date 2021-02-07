package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
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

// クライアントがgzip対応しているか確認
func gzipAcceptable(request *http.Request) bool {
	return strings.Contains(strings.Join(request.Header["Accept-Encoding"], ","), "gzip")
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

		go func() {
			defer conn.Close()
			log.Printf("Accept %v\n", conn.RemoteAddr())

			// タイムアウトを設定
			// 本来はfor文の中に置いて、リクエストが飛んでくるるたびにタイムアウトを延長すべきだが
			// タイムアウトエラー時の挙動もテストしたいので、わざとfor文の外に置いている
			err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			exitIfError(err)

			// Accept後にソケットを何度も使い回すのでループ
			for {
				// コネクションが確立されてリクエストが飛んできたらデータを読み取る
				request, err := http.ReadRequest(bufio.NewReader(conn))
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

				// 受け取ったデータを標準出力
				dump, err := httputil.DumpRequest(request, true)
				exitIfError(err)
				fmt.Printf("%s\n\n", string(dump))

				// レスポンスを書き込む
				response := http.Response{
					StatusCode: 200,
					ProtoMajor: 1,
					ProtoMinor: 1,
					Header:     make(http.Header),
				}

				if gzipAcceptable(request) {
					content := "Hello\n"

					var buffer bytes.Buffer
					writer := gzip.NewWriter(&buffer)
					_, err = io.WriteString(writer, content)
					exitIfError(err)
					writer.Close()

					response.Body = ioutil.NopCloser(&buffer)
					response.ContentLength = int64(buffer.Len())
					response.Header.Set("Content-Encoding", "gzip")
					log.Printf("raw content: %s", content)
					log.Printf("gzipped content: %+v\n", buffer.Bytes())
				} else {
					content := "Hello, World\n"
					response.Body = ioutil.NopCloser(strings.NewReader(content))
					response.ContentLength = int64(len(content))
					log.Printf("response %s\n", content)
				}
				err = response.Write(conn)
				exitIfError(err)
			}
		}()
	}
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
