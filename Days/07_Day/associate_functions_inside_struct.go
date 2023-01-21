package main

import "fmt"

// Declare a new type "PrintInfo" as a function
// that takes two arguments, a string and an int
// and returns a string
type PrintInfo func(string, int) string

// Declare a struct "Person" with fields "name", "age"
// and "Info" of type string, int and PrintInfo respectively
type Person struct {
	name string
	age  int
	Info PrintInfo
}

func main() {

	// Create a variable "result" of type "Person" and
	// initialize its fields with values
	result := Person{
		name: "Jerson Martínez",
		age:  26,
		// An anonymous function is assigned to the "Info" field
		// that takes the "name" and "age" fields of the "Person" struct and
		// concatenates them using the fmt.Sprintf function
		Info: func(name string, age int) string {
			return fmt.Sprintf("%s %d", name, age)
		},
	}

	// Print the "name" and "Info" fields of the "result" variable
	fmt.Println("Name: ", result.name)
	fmt.Println("Name and age:", result.Info(result.name, result.age))
}
