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
	clientFile := filepath.Join(os.TempDir(), "unix_domain_socket_client")
	os.Remove(clientFile)

	// Unixドメインソケットのリッスン開始
	conn, err := net.ListenPacket("unixgram", clientFile)
	exitIfError(err)
	defer conn.Close()
	log.Printf("Client is running at %s\n", conn.LocalAddr().String())

	// Unixドメインソケットへの接続
	serverFile := filepath.Join(os.TempDir(), "unix_domain_socket_server")
	unixServerAddr, err := net.ResolveUnixAddr("unixgram", serverFile)
	exitIfError(err)

	// サーバにパケットを投げる
	log.Printf("Sending to server\n")
	var serverAddr net.Addr = unixServerAddr
	_, err = conn.WriteTo([]byte("Hello from Client"), serverAddr)
	exitIfError(err)

	// サーバから送られてきたパケットを読み取る
	log.Printf("Receiving from server\n")
	buffer := make([]byte, 1500)
	length, _, err := conn.ReadFrom(buffer)
	exitIfError(err)
	fmt.Printf("Received from %v: %v\n", serverAddr, string(buffer[:length]))
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
