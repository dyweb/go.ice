package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/sleep", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("start sleeping")
		select {
		case <-request.Context().Done():
			fmt.Println("client closed connection")
		case <-time.After(2 * time.Second):
			fmt.Println("finish sleeping")
		}
		n, err := writer.Write([]byte("you can't see me if you don't wait long enough"))
		// FIXME: I was expecting to see things like broken pipe etc. after client close connection
		if err != nil {
			fmt.Printf("err %v\n", err)
		} else {
			fmt.Printf("%d bytes written", n)
		}
		if flusher, ok := writer.(http.Flusher); ok {
			flusher.Flush()
		}
	})
	addr := ":9998"
	srv := &http.Server{Addr: addr, Handler: mux}
	fmt.Printf("listen on %s\n", addr)
	fmt.Printf("visist %s/sleep and close the tab before it show something, the terminal should have some error \n", addr)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("server stoped due to %v", err)
	}
}
