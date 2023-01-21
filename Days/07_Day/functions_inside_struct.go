package main

import "fmt"

// Define la estructura "Person" con dos campos, "name" y "age"
type Person struct {
	name string
	age  int
}

// Define un método "PrintInfo" asociado a la estructura "Person"
// El método recibe una variable de tipo "Person" y utiliza los campos de esa variable para imprimir en pantalla
func (p Person) PrintInfo() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}

func main() {
	// Crea una variable "p" de tipo "Person" y asigna valores a los campos "name" y "age"
	p := Person{name: "Jerson Martínez", age: 26}
	// Llama al método "PrintInfo" de la variable "p"
	p.PrintInfo()
}
