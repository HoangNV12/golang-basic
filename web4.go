package main

import (
	"fmt"
	"net/http"
)

func liteHandler4(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Đi làm là phải chuyên nghiệp, dù rất nhớ nhà")
}
func main() {
	http.HandleFunc("/user", liteHandler4)
	http.ListenAndServe(":8080", nil)
}
