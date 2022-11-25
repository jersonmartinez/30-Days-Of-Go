package main

import (
	"fmt"
)

func main() {
	// Create a map
	var a = map[string]string{
		"Username": "jersonmartinez",
		"Name":     "Jerson Antonio",
		"Lastname": "Martínez Moreno",
	}

	fmt.Printf("=> Mapa inicial\n")
	fmt.Printf("Your Username is: " + a["Username"] + "\n")
	fmt.Printf("Your Name is: " + a["Name"] + "\n")
	fmt.Printf("Your Lastname is: " + a["Lastname"] + "\n\n")

	// Change values of the map
	a["Username"] = "CoreStack"
	a["Name"] = "Antonio"
	a["Lastname"] = "Moreno"

	fmt.Printf("=> Mapa modificado\n")
	fmt.Printf("Your Username is: " + a["Username"] + "\n")
	fmt.Printf("Your Name is: " + a["Name"] + "\n")
	fmt.Printf("Your Lastname is: " + a["Lastname"])
}
