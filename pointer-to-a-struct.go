package main

import "fmt"

type employee struct {
	name string
	empid int
}
func main()  {
	emp := employee{
		name:  "Nguyen Van Hoang",
		empid: 9999,
	}
	emp2 := emp
	fmt.Println(emp2)
	pts := &emp
	var pts2 *employee = &emp
	pts2.name = "Nguyen Thi vợ"
	// cach in ra giua mot va hai la giong nhau
	fmt.Println(pts.name)
	fmt.Println((*pts2).name)
	fmt.Println(emp2.name)
}