package main

import "fmt"

func ChangeValueByReference(x *int) {
	fmt.Println("Địa chỉ vùng nhớ con trỏ trỏ tới: ", &(*x))
	fmt.Println("Địa chỉ con trỏ", &x)
	*x = 60
}
func main() {
	var x int = 10
	var pointer *int = &x
	fmt.Println("Địa chỉ biến pointer", &(pointer))
	fmt.Println("Địa chỉ biến x trong hàm m", &x)
	fmt.Printf("Before Function Call, value of X is = %d\n", x)
	ChangeValueByReference(pointer)
	fmt.Printf("After Function Call, value of X is = %d\n", x)
}