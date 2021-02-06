package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter example")
}

func main() {
	fmt.Println("start server")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("end server")
}
