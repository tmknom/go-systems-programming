package main

import (
	"github.com/pkg/errors"
	"log"
	"net"
)

func main() {
	// 接続先の設定
	address, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
	exitIfError(err)
	log.Printf("Listen tick server: %s\n", address.String())

	// マルチキャストのリッスン開始
	listener, err := net.ListenMulticastUDP("udp", nil, address)
	defer listener.Close()

	// 送られてきたパケットを取り出す
	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := listener.ReadFrom(buffer)
		exitIfError(err)
		log.Printf("Now: %+v, by server %+v\n", string(buffer[:length]), remoteAddress)
	}
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
