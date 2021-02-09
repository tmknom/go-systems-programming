package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"syscall"
)

func main() {
	// 監視対象のファイルディスクリプタを取得
	fd, err := syscall.Open("./test", syscall.O_RDONLY, 0)
	exitIfError(err)

	// 監視したいイベントの構造体を作成
	ev1 := syscall.Kevent_t{
		Ident:  uint64(fd),
		Filter: syscall.EVFILT_VNODE,
		Flags:  syscall.EV_ADD | syscall.EV_ENABLE | syscall.EV_ONESHOT,
		Fflags: syscall.NOTE_DELETE | syscall.NOTE_WRITE,
		Data:   0,
		Udata:  nil,
	}

	// イベント待ちの無限ループ
	kq, err := syscall.Kqueue()
	exitIfError(err)
	for {
		events := make([]syscall.Kevent_t, 10)
		nev, err := syscall.Kevent(kq, []syscall.Kevent_t{ev1}, events, nil)
		exitIfError(err)
		for i := 0; i < nev; i++ {
			fmt.Printf("Event [%d] -> %+v\n", i, events[i])
		}
	}
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
