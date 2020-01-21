//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func hello() {
//	fmt.Println("hello world goroutine")
//}
//
//func main() {
//	go hello()
//	time.Sleep(1 * time.Second)
//	fmt.Println("main function")
//}
//ex2:
package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}
func main() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println()
	fmt.Println("main terminated")
}