package main

import "fmt"
/**
	trong golang cho phep khai bao method cung ten khac kieu receiver la duoc, con function thi ko
duoc phep
*/
type student struct {
	name string
	branch string
}

type teacher1 struct {
	language string
	marks int
}

func(s student) show() {
	fmt.Println("Name of student: ", s.name)
	fmt.Println("Branch: ", s.branch)
}
func (t teacher1) show()  {
	fmt.Println("Language:", t.language)
	fmt.Println("Student Marks:", t.marks)
}
func main() {
	val1 := student{
		name:   "Hoang",
		branch: "fpt",
	}
	val2 := teacher1{"java", 50}
	val1.show()
	val2.show()
}