package main

import (
	"fmt"
	"strings"
)
type FullNameByNational func(firstName, MiddleName, lastName, nationality string) string

type Person struct {
	FirstName string
	LastName string
	MiddleName string
	Nationality string
	FullName FullNameByNational
}

func main() {
	var National = [...] string {"VN", "UK"}
	p1 := Person{
		FirstName: "Hoang",
		LastName:  "Nguyen",
		MiddleName: "Van",
		Nationality: National[0],
		FullName: func(firstName, middleName, lastName, nationality string) string {
			//if nationality == "VN" {
			if strings.EqualFold(nationality, "vn") {
				return  lastName + " " + middleName + " " + firstName
			}
			return firstName + " " + middleName + " " + lastName
		},
	}
	fmt.Println(p1.FullName(p1.FirstName, p1.MiddleName, p1.LastName, p1.Nationality))
}
