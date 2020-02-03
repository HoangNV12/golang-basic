package main

import "fmt"
// neu ca hai co cung ten filed giong nhau thi se ko the truy cap truc tiep duoc (se ko la promoted field)
//struct
type details struct {
	name string
	age int
	gender string
	test string
}
// nested struct
type student1 struct {
	branch string
	year int
	test string
	details
}

func main() {
	values := student1{
		branch:  "CSE",
		year:    2020,
		test: "In student",
		details: details{
			name:   "Sumit",
			age:    30,
			gender: "Male",
			test: "in details",
		},
	}
	// Promoted fields of the student structure
	fmt.Println("Name: ", values.name)
	fmt.Println("Age: ", values.age)
	fmt.Println("Gender: ", values.gender)
	fmt.Println("Test: ", values.test)
	fmt.Println("Test: ", values.details.test)

	// Normal fields of
	// the student structure
	fmt.Println("Year: ", values.year)
	fmt.Println("Branch : ", values.branch)
}
