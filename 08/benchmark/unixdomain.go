package main

import (
	"bufio"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
)

func UnixDomainSocketStreamServer() {
	// Unixドメインソケットの準備
	socketFile := filepath.Join(os.TempDir(), "unix_domain_socket_bench")
	os.Remove(socketFile)

	// Unixドメインソケットのリッスン開始
	listener, err := net.Listen("unix", socketFile)
	exitIfError(err)
	defer listener.Close()

	for {
		// クライアントからのコネクションを待つ
		conn, err := listener.Accept()
		exitIfError(err)

		go func() {
			// コネクションが確立されてリクエストが飛んできたらデータを読み取る
			request, err := http.ReadRequest(bufio.NewReader(conn))
			exitIfError(err)

			// 受け取ったデータを読み捨てる
			_, err = httputil.DumpRequest(request, true)
			exitIfError(err)

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
