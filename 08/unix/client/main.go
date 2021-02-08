package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
)

func main() {
	// Unixドメインソケットへの接続
	socketFile := filepath.Join(os.TempDir(), "unix_domain_socket_sample")
	conn, err := net.Dial("unix", socketFile)
	exitIfError(err)

	// リクエストの生成
	request, err := http.NewRequest(
		"GET",
		"http://localhost:8888",
		nil,
	)
	exitIfError(err)

	// リクエストを送信
	err = request.Write(conn)
	exitIfError(err)

	// 受信したレスポンスの組み立て
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	exitIfError(err)

	// レスポンスを標準出力
	dump, err := httputil.DumpResponse(response, true)
	exitIfError(err)
	fmt.Printf(string(dump))
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
