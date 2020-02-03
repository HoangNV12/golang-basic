package main

import "fmt"

func myFun(a interface{}) {
	val := a.(string)
	fmt.Println("value: ", val)
}
func myFun1(a interface{}) {
	fmt.Println(a)
}
func myFun2(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("Type: int,, value: ", a.(int))
	case string:
		fmt.Println("\nType: string, Value: ", a.(string))
	case float64:
		fmt.Println("\nType: float64, Value: ", a.(float64))
	default:
		fmt.Println("\nType not found")
	}
}
func main() {
	var val interface{} = "GeeksforGeeks"
	myFun(val)
	myFun1("hoang")
	myFun1(34.78)
	myFun2("GeeksforGeeks")
	myFun2(67.9)
	myFun2(true)
}
