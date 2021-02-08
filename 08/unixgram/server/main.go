package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	// Unixドメインソケットの準備
	socketFile := filepath.Join(os.TempDir(), "unix_domain_socket_server")
	os.Remove(socketFile)

	// Unixドメインソケットのリッスン開始
	conn, err := net.ListenPacket("unixgram", socketFile)
	exitIfError(err)
	defer conn.Close()
	log.Printf("Server is running at %s\n", conn.LocalAddr().String())

	buffer := make([]byte, 1500)
	for {
		// クライアントから送られてきたパケットを読み取る
		length, remoteAddress, err := conn.ReadFrom(buffer)
		exitIfError(err)
		fmt.Printf("Received from %v: %v\n", remoteAddress, string(buffer[:length]))

		// クライアントにパケットを投げる
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)
		exitIfError(err)
	}
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
