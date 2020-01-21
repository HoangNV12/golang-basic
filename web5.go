package main

import (
	"fmt"
	"net/http"
	"time"
)

func liteHandler5(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Tết đến rồi, về nhà sum họp thôi")
}
func main() {
	http.HandleFunc("/welcome", liteHandler5)
	server := &http.Server{
		Addr : ":8080",
		ReadTimeout:10*time.Second,
		WriteTimeout:10*time.Second,
		MaxHeaderBytes:1<<20,
	}
	server.ListenAndServe()
}
