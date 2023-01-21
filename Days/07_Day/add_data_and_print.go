package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{name: "Jerson Martínez", age: 26}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}
