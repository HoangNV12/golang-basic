package main

import (
	"fmt"
	"net/http"
)

type liteHandler struct {
	message string
}
func(m *liteHandler) ServeHTTP(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, m.message)
}
func main() {
	hdl := &liteHandler{"Chào mừng đến lập trình Go cho web"}
	fmt.Println("Listening....")
	http.ListenAndServe(":9996", hdl)
}