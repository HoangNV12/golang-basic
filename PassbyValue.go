package main

import "fmt"

func changeValue(x int) {
	fmt.Println("Address of x in changeValue function: ", &x)
	x = 50
}

func main() {
	x := 20
	fmt.Println("Address of x in main function: ", &x)

	fmt.Printf("Before Function call, value of X is = %d\n", x)
	changeValue(x)
	fmt.Printf("After Function call, value of x is = %d\n", x)
}

