package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")

	const filename = "log.txt"
	file, _ := os.Create(filename)
	defer file.Close()

	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()

	writer := io.MultiWriter(file, gzipWriter)

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	encoder.Encode(map[string]string{
		"hello": "world",
	})

	gzipWriter.Flush()
}

func main() {
	fmt.Println("start server")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("end server")
}
