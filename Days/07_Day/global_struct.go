package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var p Person
	p.name = "Jerson Martínez"
	p.age = 26
	fmt.Println("Información de persona en la función main:")
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)

	fmt.Println("\nInformación de persona en la función printPerson:")
	printPerson(p)
}

func printPerson(person Person) {
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)
}
