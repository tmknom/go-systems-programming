package main

import (
	"github.com/pkg/errors"
	"gopkg.in/fsnotify.v1"
	"log"
)

func main() {
	counter := 0
	watcher, err := fsnotify.NewWatcher()
	exitIfError(err)
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Printf("%+v\n", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("created file:", event.Name)
					counter++
				} else if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					counter++
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("removed file:", event.Name)
					counter++
				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
					log.Println("renamed file:", event.Name)
					counter++
				} else if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					log.Println("chmod file:", event.Name)
					counter++
				}
			case err := <-watcher.Errors:
				log.Printf("%+v\n", errors.WithStack(err))
			}
			if counter > 3 {
				done <- true
			}
		}
	}()
	err = watcher.Add(".")
	exitIfError(err)
	<-done
}

func exitIfError(err error) {
	if err != nil {
		log.Fatalf("%+v\n", errors.WithStack(err))
	}
}
