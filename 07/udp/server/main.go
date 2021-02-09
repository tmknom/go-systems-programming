package main

import (
	"github.com/pkg/errors"
	"log"
	"net"
	"time"
)

func main() {
	// パケットのリッスン開始
	conn, err := net.ListenPacket("udp", "localhost:8888")
	exitIfError(err)
	defer conn.Close()
	log.Printf("Server is running at %s\n", conn.LocalAddr().String())

	buffer := make([]byte, 1500)
	for {
		// 送られてきたパケットを取り出す
		length, remoteAddress, err := conn.ReadFrom(buffer)
		exitIfError(err)
		log.Printf("Recieved from %+v %+v\n", remoteAddress, string(buffer[:length]))

		time.Sleep(5 * time.Second)

		// クライアントにパケットを投げる
		_, err = conn.WriteTo([]byte("Hello from server"), remoteAddress)
		exitIfError(err)
	}
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
