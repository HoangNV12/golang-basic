package main

import (
	"fmt"
	"sync"
)

//var x int = 0
//func tangX() {
//	x = x + 1
//	wg1.Done()
//}
//var wg1 sync.WaitGroup
//func main() {
//	wg1.Add(1000)
//	for i := 0; i < 1000; i++ {
//		go tangX()
//	}
//	wg1.Wait()
//	time.Sleep(2 * time.Second)
//	fmt.Println(x)
//}

//var x int = 0
//func increment(wg *sync.WaitGroup, m *sync.Mutex) {
//	m.Lock()
//	x = x + 1
//	m.Unlock()
//	wg.Done()
//}
//func main() {
//	var w sync.WaitGroup
//	var m sync.Mutex
//	for i := 0; i < 1000; i++ {
//		w.Add(1)
//		go increment(&w, &m)
//	}
//	w.Wait()
//	fmt.Println("final value of x:", x)
//}

var x int
func increment(w *sync.WaitGroup, b chan bool) {
	b <- true
	x += 1
	<- b
	w.Done()
}
func main() {
	var w sync.WaitGroup
	b := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, b)
	}
	w.Wait()
	fmt.Println(x)
}