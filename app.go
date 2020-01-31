package main

import "fmt"

type person struct {
	name string
	age int
}
type teacher struct {
	bang string
	soNamKinhNghiem int
	baseInfo person
}

func main()  {
	t1 := teacher{
		bang:            "DH Dai Nam",
		soNamKinhNghiem: 4,
		baseInfo:        person{
			name: "Nguyen Van Hoang",
			age: 30,
		},
	}
	//in thong tin
	fmt.Println(t1.baseInfo.name)
	fmt.Println(t1.bang)

}