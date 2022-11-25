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

	// Looping through the map
	for key, value := range User {
		fmt.Printf("The %s is: %s \n", key, value)
	}

}
