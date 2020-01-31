package main

import "fmt"

func ChuyenHamAnonyMousVaoFunc(i func(a string) string) {
	fmt.Println(i("Hoang"))
}
func TraVeHamAnonymous() func(i string) string {
	return func(i string) string {
		return i + " is a blogger"
	}
}
func main()  {
	ChuyenHamAnonyMousVaoFunc(func(name string) string {
		return "Welcome " + name
	})
	greeting := func(name string) string{
		return "Hello " + name
	}
	ChuyenHamAnonyMousVaoFunc(greeting)
	introduction := TraVeHamAnonymous()
	fmt.Println(introduction("Hoang"))
}