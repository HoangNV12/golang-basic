package main

import "fmt"

func Exam4(i func(p, q string) string) {
	//bien i co kieu du lieu la func, func co hai doi so truyen vao,  va func return ve string
	fmt.Println(i("day la p", "day la q"))
}
func Exam5() func(i, j string) string {
	myF := func(i, j string) string {
		return i + j + "GeeksforGeeks"
	}
	return myF
}
func main() {
	//Ham an danh

	//ex1:
	func() {
		fmt.Println("Welcome to golang")
	}()

	//ex2: gan ham an danh vao bien, sau do su dung
	var value func()
	value = func() {
		fmt.Println("Vi du gan ham an danh vao bien")
	}
	value()
	//ex3: chuyen doi so trong ham an danh
	func(ele string) {
		fmt.Println(ele)
	}("Tôi là đối số truyền vào")
	//ex4: chuyền hàm ẩn danh như là một đối số

	value1 := func(p, q string) string {
		return p + " " +  q + " la hai gia tri chuyen vao"
	}
	Exam4(value1)
	//ex5: tra ve mot ham an danh
	value2 := Exam5()
	fmt.Println(value2("Welcome", "to"))
}
