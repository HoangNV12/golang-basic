package main

import "fmt"

//Ex1:
//func ChangeValueByReference(x *int) {
//	fmt.Println("Địa chỉ vùng nhớ con trỏ trỏ tới: ", &(*x))
//	fmt.Println("Địa chỉ con trỏ", &x)
//	*x = 60
//}
//func main() {
//	var x int = 10
//	var pointer *int = &x
//	fmt.Println("Địa chỉ biến pointer", &(pointer))
//	fmt.Println("Địa chỉ biến x trong hàm m", &x)
//	fmt.Printf("Before Function Call, value of X is = %d\n", x)
//	ChangeValueByReference(pointer)
//	fmt.Printf("After Function Call, value of X is = %d\n", x)
//}
//Ex2:
func swap(x, y *int) int {
	var tmp int

	tmp = *x
	*y = *x
	*y = tmp
	return tmp
}
func main() {
	x := 30
	y := 40
	fmt.Printf("Values Before Function Call\n")
	fmt.Printf("f = %d and s = %d\n", x, y)
	swap(&x, &y)

	fmt.Printf("\nValues After Function Call\n")
	fmt.Printf("f = %d and s = %d", x, y)
}