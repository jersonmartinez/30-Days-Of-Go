package main

import (
	"fmt"
)

func main() {
	// Create a map
	var User = map[string]string{
		"Username": "jersonmartinez",
		"Name":     "Jerson Antonio",
		"Lastname": "Martínez Moreno",
	}

	// Add element
	User["Age"] = "26"

	// Add other element
	User["Career"] = "DevOps Engineer"

	fmt.Printf("Your Username is: " + User["Username"] + "\n")
	fmt.Printf("Your Name is: " + User["Name"] + "\n")
	fmt.Printf("Your Lastname is: " + User["Lastname"] + "\n")

	// New elements
	fmt.Printf("Your Age is: " + User["Age"] + "\n")
	fmt.Printf("Your Career is: " + User["Career"] + "\n\n")

}
