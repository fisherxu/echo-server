package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println()
	r := make(chan string)
	echo := func(response http.ResponseWriter, request *http.Request) {
		b, _ := io.ReadAll(request.Body)
		fmt.Println(string(b))
		if _, err := response.Write([]byte("Hello World")); err != nil {
			log.Fatalf("Echo server write failed. reason: %s", err.Error())
		}
	}
	url := func(response http.ResponseWriter, request *http.Request) {
		b, _ := io.ReadAll(request.Body)
		var buff bytes.Buffer
		buff.WriteString("Reply from server: ")
		buff.Write(b)
		buff.WriteString(" Header of the message: [user]: " + request.Header.Get("user") +
			", [passwd]: " + request.Header.Get("passwd"))
		if _, err := response.Write(buff.Bytes()); err != nil {
			log.Fatalf("Echo server write failed. reason: %s", err.Error())
		}
		r <- buff.String()
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/cpu-stats", echo)
	mux.HandleFunc("/url", url)
	server := &http.Server{Addr: "127.0.0.1:8080", Handler: mux}

	server.ListenAndServe()

}
