package main

import (
	"fmt"
	"log"
	"net/http"
)

type liteHandler1 struct {
	message string
}
func(m *liteHandler1) ServeHTTP(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, m.message)
}
func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./html"))
	mux.Handle("/", fs)
	hdl := &liteHandler1{"Chào mừng đến lập trình Go cho web"}
	mux.Handle("/welcome", hdl)
	fmt.Println("Listening....")
	log.Fatal(http.ListenAndServe(":8080", mux))
}