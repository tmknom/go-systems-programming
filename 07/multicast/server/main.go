package main

import (
	"github.com/pkg/errors"
	"log"
	"net"
	"time"
)

const interval = 10 * time.Second

func main() {
	// 送信先の設定
	conn, err := net.Dial("udp", "224.0.0.1:9999")
	exitIfError(err)
	defer conn.Close()
	log.Printf("Start tick server at %s\n", conn.LocalAddr().String())

	// 現実時刻のn*10秒になるよう、少しだけ待つ
	start := time.Now()
	wait := start.Round(interval).Add(interval).Sub(start)
	time.Sleep(wait)

	// intervalの間隔で定期実行する
	ticker := time.Tick(interval)
	for now := range ticker {
		tick := now.String()
		conn.Write([]byte(tick))
		log.Printf("Tick: %s\n", tick)
	}
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
