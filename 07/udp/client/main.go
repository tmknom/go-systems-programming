package main

import (
	"github.com/pkg/errors"
	"log"
	"net"
)

func main() {
	// 接続先の設定
	conn, err := net.Dial("udp4", "localhost:8888")
	exitIfError(err)
	defer conn.Close()
	log.Printf("Sending to %s server\n", conn.RemoteAddr().String())

	// サーバへパケットを投げる
	_, err = conn.Write([]byte("Hello from Client"))
	exitIfError(err)
	_, err = conn.Write([]byte("Hello from God"))
	exitIfError(err)

	// サーバからパケットを受け取る
	log.Printf("Recieving from server...\n")
	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)
	exitIfError(err)
	log.Printf("Recieved from %+v %+v\n", conn.RemoteAddr().String(), string(buffer[:length]))
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
