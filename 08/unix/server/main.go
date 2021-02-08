package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Unixドメインソケットの準備
	socketFile := filepath.Join(os.TempDir(), "unix_domain_socket_sample")
	os.Remove(socketFile)

	// Unixドメインソケットのリッスン開始
	listener, err := net.Listen("unix", socketFile)
	exitIfError(err)
	defer listener.Close()
	log.Printf("Server is running at %s\n", listener.Addr().String())

	for {
		// クライアントからのコネクションを待つ
		conn, err := listener.Accept()
		exitIfError(err)

		go func() {
			log.Printf("Accept %v\n", conn.RemoteAddr())

			// コネクションが確立されてリクエストが飛んできたらデータを読み取る
			request, err := http.ReadRequest(bufio.NewReader(conn))
			exitIfError(err)

			// 受け取ったデータを標準出力
			dump, err := httputil.DumpRequest(request, true)
			exitIfError(err)
			fmt.Printf(string(dump))

			// レスポンスを書き込む
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World\n")),
			}
			err = response.Write(conn)
			exitIfError(err)
			conn.Close()
		}()
	}
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
