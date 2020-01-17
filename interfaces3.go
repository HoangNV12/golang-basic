package main

type Describer interface {
	Describe()
}

func main() {
	// chuong trinh se be loi vi co trinh truy cap vao null
	var d1 Describer
	d1.Describe()
}