package main

import (
	"fmt"
	"reflect"
)

/**
Trong go co the dung dau == or ham DeeplyEqual (reflect package) để so sánh hai struct
Structs là các kiểu giá trị và có thể so sánh được nếu mỗi trường của chúng có thể so sánh được.
Hai biến cấu trúc được coi là bằng nhau nếu các trường tương ứng của chúng bằng nhau.
Các biến cấu trúc không thể so sánh nếu chúng chứa các trường không thể so sánh được (vd: map, slice, array...)
dùng dấu bằng sẽ bị lỗi, nhưng dùng hàm  DeepEqual thì được
*/
type Author3 struct {
	name string
	language string
	Particles int
}
type Image struct {
	data map[int]int
}
func main() {
	img1 := Image{data: map[int]int{
		0 : 155,
	}}
	img2 := Image{data: map[int]int{
		0 : 155,
	}}
	// chuong trinh se bi loi luc bien dich
	//if img1 == img2 {
	//	fmt.Println("Variable img1 equal img2")
	//} else {
	//	fmt.Println("Variable img1 not equal img2")
	//}

	fmt.Println("Is a1 equal to a2: ", reflect.DeepEqual(img1, img2))
	a1 := Author3{
		name:      "Hoang",
		language:  "java",
		Particles: 35,
	}

	a2 := Author3{
		name:      "Hoang",
		language:  "java",
		//Particles: 35,
	}
	if a1 == a2 {
		fmt.Println("Variable a1 equal a2")
	} else {
		fmt.Println("Variable a1 not equal a2")
	}
	fmt.Println("Is a1 equal to a2: ", reflect.DeepEqual(a1, a2))
}
