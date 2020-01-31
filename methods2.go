package main

import (
	"fmt"
	"reflect"
)

type data int
type haha int
func (d1 data) multiply(d2 data) data {
	return d1 * d2
}
func (d1 haha) testM(d2 haha) haha {
	return d1 * d2
}
func main() {
	value1 := data(300)
	value2 := data(20)
	fmt.Println(reflect.TypeOf(value1).String())
	res := value1.multiply(value2)
	fmt.Println("Final result: ", res)

	value3 := haha(300)
	value4 := haha(30)
	res1 := value3.testM(value4)
	fmt.Println("Final result: ", res1)
}
