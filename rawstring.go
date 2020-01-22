package main

import (
	"fmt"
	"time"
)

func main() {
	str := ` dong 1
			Dong 2
			dong 3
			don 4
			"dong 5"
			'dong 6'
			`
	fmt.Println(str)
	str1 := fmt.Sprintf("at %v, %s", time.Now(), "Now")
	fmt.Println(str1)
	var myArr [5]int
	myArr[0] = 1
	myArr[1] = 2
	myArr[2] = 6
	myArr[3] = 7
	myArr[4] = 5
	arr := myArr
	fmt.Println(&myArr)
	fmt.Println(&arr)
}
