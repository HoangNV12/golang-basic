package main

import "fmt"

type author struct {
	name string
	branch string
	particles int
	salary int
}

func(a *author) show() {
	fmt.Println("Author's name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}
func showInfo(a author) {
	fmt.Println("Author's name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}

func main() {
	res := author{
		name:      "Hoang",
		branch:    "Javascript",
		particles: 200,
		salary:    3400,
	}
	res.show()
	res2 := &res
	res2.show()
	//showInfo(&res)
}