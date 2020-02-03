// Golang program to illustrate the 
// concept of the promoted methods 
package main

import "fmt"

// Structure 
type xdetails2 struct {

	// Fields of the 
	// details structure 
	name string
	age	 int
	gender string
	psalary int
}

// Method 1 
func (e xemployee2) promotmethod(tarticle int) int {
	return e.particle * tarticle
}

// Nested structure 
type xemployee2 struct {
	post	 string
	particle int
	eid	 int
	xdetails2
}

// Method 2 
func (d xdetails2) promotmethod(tsalary int) int {
	return d.psalary * tsalary
}

// Main method 
func main() {

	// Initializing the fields of 
	// the employee structure 
	values := xemployee2{
		post:	 "HR",
		eid:	 4567,
		particle: 5,
		xdetails2: xdetails2{

			name: "Sumit",
			age:	 28,
			gender: "Male",
			psalary: 890,
		},
	}

	// Promoted fields of 
	// the employee structure 
	fmt.Println("Name: ", values.name)
	fmt.Println("Age: ", values.age)
	fmt.Println("Gender: ", values.gender)
	fmt.Println("Per day salary: ", values.psalary)

	// Promoted method of 
	// the employee structure 
	fmt.Println("Total Salary: ", values.xdetails2.promotmethod(30))

	// Normal fields of 
	// the employee structure 
	fmt.Println("Post: ", values.post)
	fmt.Println("Employee id: ", values.eid)
	fmt.Println("Total Articles: ", values.promotmethod(30))
} 
