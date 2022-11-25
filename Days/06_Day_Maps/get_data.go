package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]string)
	a["Username"] = "jersonmartinez"
	a["Name"] = "Jerson Antonio"
	a["Lastname"] = "Martínez Moreno"

	fmt.Printf("Your Username is: " + a["Username"] + "\n")
	fmt.Printf("Your Name is: " + a["Name"] + "\n")
	fmt.Printf("Your Lastname is: " + a["Lastname"])
}
