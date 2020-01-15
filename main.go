package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "Golang"

	value := <- messages

	fmt.Println(value)
}
