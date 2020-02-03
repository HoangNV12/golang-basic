package main

import (
	"fmt"
	"math"
)

/**
trong go, ko the tao mot mot instance of interface, nhung co the tao mot bien co kieu la interface va gan
mot gia tri co kieu cu the cho bien do (kieu du lieu phai implement tat ca cac method trong interface)
de implement interface thi golang ko ho tro mot tu khoa nao ca, ma no duoc thuc hien ngam dinh (khi tat ca cac method
trong interface duoc trien khai)
*/

//dinh nghia mot interface chua hai phuong thuc tinh dien tich va the tich
type tank interface {
	//methods
	Tarea() float64
	Volume() float64
}
type tanks interface {
	Tarea() float64
}
type myvalue struct {
	radius float64
	height float64
}

// trien khai (thuc thi) cac phuong thuc cua interface (Implementing methods of the tank inteface)

func (m myvalue) Tarea() float64 {
	return 2*m.radius*m.height + 2*math.Pi*math.Pow(m.radius, 2)
}

func (m myvalue) Volume() float64 {
	return 3.14 * m.radius * m.radius * m.height
}

func main() {
	var t tank
	t = myvalue{
		radius: 10,
		height: 14,
	}
	var t1 tanks
	t1 = myvalue{
		radius: 10,
		height: 14,
	}
	fmt.Println(t1.Tarea())
	fmt.Println("Area of tank:", t.Tarea())
	fmt.Println("Volume of tank: ", t.Volume())
}
