package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
	}

	p := Person{name: "Jerson Martínez", age: 26}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}
