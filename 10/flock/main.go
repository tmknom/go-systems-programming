package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"sync"
	"syscall"
	"time"
)

type FileLock struct {
	l  sync.Mutex
	fd int
}

func NewFileLock(filename string) *FileLock {
	if filename == "" {
		panic("filename needed")
	}
	fd, err := syscall.Open(filename, syscall.O_CREAT|syscall.O_RDONLY, 0750)
	exitIfError(err)
	return &FileLock{fd: fd}
}

func (m *FileLock) Lock() {
	m.l.Lock()
	err := syscall.Flock(m.fd, syscall.LOCK_EX)
	exitIfError(err)
}

func (m *FileLock) UnLock() {
	err := syscall.Flock(m.fd, syscall.LOCK_UN)
	exitIfError(err)
	m.l.Unlock()
}

func main() {
	l := NewFileLock("main.go")
	fmt.Printf("try locking...\n")
	l.Lock()
	fmt.Printf("locked!\n")
	time.Sleep(10 * time.Second)
	l.UnLock()
	fmt.Printf("unlock\n")
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
