package main

import "fmt"

type author1 struct {
	name string
	branch string
}

func(a *author1) show_2() {
	a.name = "name in show 2"
	fmt.Println("Author's name(Before): ", a.name)
}
func main()  {
	res := author1{
		name:   "Sona",
		branch: "CSE",
	}
	fmt.Println("name in main function (Before)", res.name)
	res.show_2()
	fmt.Println("name in main function (After)", res.name)
}