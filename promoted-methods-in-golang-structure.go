package main

import "fmt"

//struct
type details1 struct {
	name string
	age int
	gender string
	psalary int
}
// nested struct
type employee1 struct {
	post string
	eid int
	details1
}
func (d details1) promotmethod(tsalary int) int {
	return d.psalary * tsalary
}
func main() {
	values := employee1{
		post:     "HR",
		eid:      4567,
		details1: details1{
			name:    "sumit",
			age:     28,
			gender:  "Male",
			psalary: 890,
		},
	}
	//promoted fields of the employee1 struct
	fmt.Println("Name:", values.name)
	//promoted method of the employee1 struct
	fmt.Println("Total salary:", values.promotmethod(30))

}