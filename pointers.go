package main

import (
	"fmt"
)
type Teacher struct {
	name string
}
type Teacher1 struct {
	name string
	welcome func(interface{}) string
}
func GetName() *Teacher {
	return &Teacher{"Nguyen Van Hoang"}
}
func printAny(string) {
	fmt.Println()
}
func main() {
	printAny("abc")
	i := 10
	fmt.Println(i)
	fmt.Println(&i)
	var p1 *int
	fmt.Println(p1)
	p1 = &i
	p2 := &i
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(&p1)
	fmt.Println(*p1)
	fmt.Printf("%T\n", &p1)
	if p1 == p2 {
		fmt.Println("p1 equal p2")
	} else {
		fmt.Println("p1 not equal p2")
	}
	var t1 *Teacher = GetName()
	t2 := GetName()
	fmt.Println(t1.name)
	fmt.Println(t2.name)
	fmt.Println(t1)
	fmt.Println(t2)
	if t1 != t2 {
		fmt.Println("t1 not equal t2")
	} else {
		fmt.Println("t1 equal t2")
	}
	var t3 Teacher1
	t3.name = "hoang"
	fmt.Println(t3.name)
	t3.welcome = func(message interface{}) string {
		return "Hello," + message.(string)
	}
	var t4 Teacher1
	t4.name = "Manh"
	t4.welcome = func(i interface{}) string {
		return "welcome " + fmt.Sprintf("%v",i)
	}
	fmt.Println(t3.welcome("hoang"))
	fmt.Println(t4.welcome(t4.name))
}