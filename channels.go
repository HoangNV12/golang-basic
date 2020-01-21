//package main
//
//import (
//	"fmt"
//)
//
//func hello(done chan bool) {
//	fmt.Println("Hello world goroutine")
//	done <- true
//}
//func main() {
//	//exp1:
//	//var a chan int
//	//if a == nil {
//	//	fmt.Println("Channel a is nil, going to define it")
//	//	a = make(chan int)
//	//	fmt.Printf("Type of a is %T", a)
//	//}
//
//	//exp 2:
//	done := make(chan bool)
//	go hello(done)
//	<-done
//	fmt.Println("main function")
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func hello(done chan bool) {
//	fmt.Println("hello go routine is going to sleep")
//	time.Sleep(4 * time.Second)
//	fmt.Println("hello go routine awake and going to write to done")
//	done <- true
//}
//func main() {
//	done := make(chan bool)
//	fmt.Println("Main going to call hello go goroutine")
//	go hello(done)
//	<- done
//	fmt.Println("Main received data")
//}

//package main
//
//import (
//	"fmt"
//)
//
//func calcSquare(number int, squareop chan int) {
//	sum := 0
//	for number != 0 {
//		digit := number % 10
//		sum += digit * digit
//		number /= 10
//	}
//	squareop <- sum
//}
//func calcCubes(number int, cubeop chan int) {
//	sum := 0
//	for number != 0 {
//		digit := number % 10
//		sum += digit * digit * digit
//		number /= 10
//	}
//	cubeop <- sum
//}
//func main() {
//	number := 589
//	sqrch := make(chan int)
//	cubech := make(chan int)
//	go calcSquare(number, sqrch)
//	go calcCubes(number, cubech)
//	squares := <- sqrch
//	cubes := <- cubech
//	fmt.Println("Final output", squares + cubes)
//}

package main

import "fmt"

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println("Received ", v)
	}
}