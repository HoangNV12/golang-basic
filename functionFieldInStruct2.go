package main

import "fmt"

//Tao mot struct
type Author1 struct {
	name      string
	language  string
	Tarticles int
	Particles int
	Pending   func(int, int) int
}

func main() {
	//khoi tao gia tri
	result := Author1{
		name:      "Alex",
		language:  "Java",
		Tarticles: 340,
		Particles: 259,
		Pending: func(Ta int, Pa int) int {
			return Ta - Pa
		},
	}
	// Hien thi gia tri
	fmt.Println("Author's Name: ", result.name)
	fmt.Println("Language: ", result.language)
	fmt.Println("Total number of articles: ", result.Tarticles)

	fmt.Println("Total number of published articles: ",
		result.Particles)

	fmt.Println("Pending articles: ", result.Pending(result.Tarticles,
		result.Particles))
}
