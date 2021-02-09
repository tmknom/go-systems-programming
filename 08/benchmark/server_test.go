package main

import (
	"bufio"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func BenchmarkTCPServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// TCPサーバへの接続
		conn, err := net.Dial("tcp", "localhost:18888")
		exitIfError(err)

		// 通信の実行
		run(conn)
	}
}

func BenchmarkUnixDomainSocketStreamServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Unixドメインソケットサーバへの接続
		socketFile := filepath.Join(os.TempDir(), "unix_domain_socket_bench")
		conn, err := net.Dial("unix", socketFile)
		exitIfError(err)

		// 通信の実行
		run(conn)
	}
}

func run(conn net.Conn) {
	// リクエストの生成
	request, err := http.NewRequest(
		"GET",
		"http://localhost:18888",
		nil,
	)
	exitIfError(err)

	// リクエストを送信
	err = request.Write(conn)
	exitIfError(err)

	// 受信したレスポンスの組み立て
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	exitIfError(err)

	// 受け取ったデータを読み捨てる
	_, err = httputil.DumpResponse(response, true)
	exitIfError(err)

	conn.Close()
}

func TestMain(m *testing.M) {
	// init
	go UnixDomainSocketStreamServer()
	go TCPServer()
	time.Sleep(time.Second)
	// run test
	code := m.Run()
	// exit
	os.Exit(code)
}
