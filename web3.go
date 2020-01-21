package main

import (
	"fmt"
	"log"
	"net/http"
)

func liteHandler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Chào mừng bạn đến lập trình Go cho web")
}
func findUserById(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "Find user by id")
}
func findClassroom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I miss topica")
}
func main() {
	mux := http.NewServeMux()
	// ép kiểu liteHandler3 về http.HandleFunc
	hdl := http.HandlerFunc(liteHandler3)
	mux.Handle("/welcome", hdl)
	user := http.HandlerFunc(findUserById)
	mux.Handle("/user", user)
	mux.HandleFunc("/classroom", findClassroom)
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
