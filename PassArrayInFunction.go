package main

import "fmt"

/**
	// For sized array
	func function_name(variable_name [size]type){
	// Code
	}

	// For unsized array (this is slice)
	func function_name(variable_name []type){
	// Code
	}

*/

func myFunc(a [6]int, size int) int {
	var val int

	for k := 0; k < size; k++ {
		val += a[k]
	}
	return val/size
 }
 func myFunc2(a []int) int {
 	var val int
	 var size int = len(a)
	 for k := 0; k < size; k++ {
	 	val += a[k]
	 }
	 return val/size
 }

 func main() {
 	var arr = [6] int {67, 59, 29, 35, 4, 34}
 	var res int

 	res = myFunc(arr, 6)
 	fmt.Printf("Final result is: %d\n", res)

 	arr1 := []int {1,2,3,4,5,6,7}
 	res = myFunc2(arr1)
 	fmt.Printf("Final result2 is: %d\n", res)
 }