package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)
//https://hung.phamkhac.com/2016/10/bai-27-xu-ly-bieu-mau.html
var id int = 0
var locationStore = make(map[int]Location)
type Location struct {
	Name      string  `json:"name"`
	Address   string  `json:"adr"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Type      string  `json:"type"`
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var locpage string = `
			<!DOCTYPE html>
			<html>
				<head>
					<title>Địa điểm mới</title>
				</head>
				<body>
					<h2>Thêm địa điểm mới</h2>
					<form action="/locations" method="post">
						<table>
							<tr>
								<td>Tên</td>
								<td><input type="text" name="name" /></td>
							</tr>
							<tr>
								<td>Địa chỉ</td>
								<td><input type="text" name="adr"></td>
							</tr>
							<tr>
								<td>Vĩ độ</td>
								<td><input type="text" name="lat"></td>
							</tr>
							<tr>
								<td>Kinh độ</td>
								<td><input type="text" name="lon"></td>
							</tr>
							<tr>
								<td>Loại</td>
								<td>
									<select name="type">
										<option value="-">-</option>
										<option value="Ăn uống">Ăn uống</option>
										<option value="Nghỉ ngơi">Nghỉ ngơi</option>
										<option value="Đi lại">Đi lại</option>
										<option value="Giải trí">Giải trí</option>
										<option value="Tiện ích">Tiện ích</option>
									</select>
								</td>
							</tr>
							<tr>
								<td>&nbsp</td>
								<td><input type="submit" value="Thêm"></td>
							</tr>
						</table>			
					</form>
				</body>
			</html>
		`
		fmt.Fprintln(w, locpage)
	} else {
		var location Location
		r.ParseForm()
		location.Name = template.HTMLEscapeString(r.PostForm.Get("name"))

		if len(strings.TrimSpace(location.Name)) == 0 {
			fmt.Fprintf(w, "Thêm không thành công do tên địa điểm rỗng!\nThông tin các địa điểm: %v", locationStore)
			return
		}
		location.Address = r.PostForm.Get("adr")
		var err error
		location.Latitude, err = strconv.ParseFloat(r.PostForm.Get("lat"), 32)
		if err != nil {
			panic(err)
		}
		location.Longitude, err = strconv.ParseFloat(r.PostForm.Get("lon"), 32)
		if err != nil {
			panic(err)
		}
		//location.Type = r.PostForm.Get("type")
		loctype := []string{"Ăn uống", "Nghỉ ngơi", "Đi lại", "Giải trí", "Tiện ích"}
		for _, v := range loctype {
			if v == r.PostForm.Get("type") {
				location.Type = v
				break
			}
		}
		if len(location.Type) == 0 {
			fmt.Fprintf(w, "Thêm không thành công do sai loại địa điểm!\nThông tin các địa điểm: %v", locationStore)
			return
		}
		id++
		locationStore[id] = location
		fmt.Fprintf(w, "Thêm thành công!\nThông tin các địa điểm: %v", locationStore)
	}
}
func main() {
	http.HandleFunc("/", LocationHandler)
	http.HandleFunc("/locations", LocationHandler)
	http.ListenAndServe(":8080", nil)
}
