//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	ch := make(chan int, 2)
//	ch <- 5
//	ch <- 8
//	fmt.Println(cap(ch))
//	fmt.Println(len(ch))
//	fmt.Println(<- ch)
//	fmt.Println(len(ch))
//	fmt.Println(<- ch)
//}
//

//func main() {
//c := make(chan int)
//go func() {
//c <- 1
//}()
//data := <- c
//fmt.Println(data)
//}

package main

//func write(ch chan int)  {
//	for i := 0; i < 5; i ++ {
//		ch <- i
//		fmt.Println("ghi ", i, "vao kenh")
//	}
//	close(ch)
//}
//func main() {
//	ch := make(chan int, 2)
//	go write(ch)
//	//time.Sleep(2 * time.Second)
//	for v := range ch {
//		fmt.Println("doc gia tri", v, "tu kenh")
//		//time.Sleep(2*time.Second)
//	}
//}

//func process(i int, wg *sync.WaitGroup) {
//	fmt.Println("started Goroutin ", i)
//	time.Sleep(2*time.Second)
//	fmt.Printf("Goroutine %d ended\n", i)
//	wg.Done()
//}
//func main() {
//	no := 3
//	var wg sync.WaitGroup
//	for i := 0; i < no; i++ {
//		wg.Add(1)
//		go process(i, &wg)
//	}
//	wg.Wait()
//	fmt.Println("All go routines finished executing")
// }