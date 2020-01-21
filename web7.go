package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("upfile")
	if err != nil {
		fmt.Fprintln(w, "Tải file bị lỗi")
		return
	}
	defer file.Close()
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		fmt.Fprintln(w, "Đọc file lỗi")
		return
	}
	filetype := http.DetectContentType(buff)
	if filetype == "image/jpeg" || filetype == "image/jpg" || filetype == "image/png" {
		f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Fprintln(w, "Lưu file lỗi")
		}
		defer f.Close()
		io.Copy(f, file)
		fmt.Fprintln(w, "Tải file thành công:%v", handler.Header, filetype)
	} else {
		fmt.Fprintf(w, "Định dạng file không hỗ trợ: %s", filetype)
	}

}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var template string = `
			<!DOCTYPE html>
			<html>
				<head>
					<title>Example upload file in go</title>
				</head>
				<body>
					<form enctype="multipart/form-data" action="/upload" method="post">
						<input type="file" name="upfile" accept=".jpg,.png"/>
						<input type="submit" value="Tải lên" />
					</form>
				</body>
			</html>
		`
		fmt.Fprintln(w, template)
	})
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}
