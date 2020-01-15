package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func printCounts(label string, count chan int) {
	defer wg.Done()
	for {
		val, ok := <- count
		if !ok {
			fmt.Println("Channel was closed")
			return
		}
		fmt.Printf("Count: %d received from %s \n", val, label)
		if val == 10 {
			fmt.Printf("Channel closed from %s \n", label)
			close(count)
			return
		}
		val++
		count <- val
	}
}
func main() {
	count := make(chan int)
	wg.Add(2)
	fmt.Println("Start Goroutines")
	go printCounts("A", count)
	go printCounts("B", count)
	count <- 1
	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("\nTerminating program")
}
