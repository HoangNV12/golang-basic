package main

import "fmt"

func main() {
	Element := struct {
		name string
		branch string
		language string
		Particles int
	}{
		// khoi tao gia tri o day, hoac khoi tao gia tri o ngoai
		name: "Pikachu",
		branch: "ECE",
		language: "C++",
	}
	Element.Particles = 48
	fmt.Println(Element)
}
