package main 

import (
	"fmt"
)
type SinhVien struct {
	FirstName string
	LastName string
	MiddleName string
}
func main() {
	fmt.Println("test")
	sv := SinhVien {
		FirstName 	: "Hoang",
		MiddleName 	: "Van",
		LastName 	: "Nguyen",
	}

	fmt.Println(sv)

	st := struct {
		name string
	} {
			name: "Day la struct anonymous co phai k?",
	}
	fmt.Println(st)
}