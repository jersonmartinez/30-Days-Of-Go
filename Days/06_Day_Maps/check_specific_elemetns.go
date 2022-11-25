package main

import "fmt"

func main() {
	// Create a map
	var User = map[string]string{
		"Username": "jersonmartinez",
		"Name":     "Jerson Antonio",
		"Lastname": "Martínez Moreno",
		"Age":      "",
	}

	// Check if exists key and its value
	k1, v1 := User["Username"]

	// Check if non-exists key and its value
	k2, v2 := User["Phone"]

	// Check if exists key and its value
	k3, v3 := User["Age"]

	// Only check if exists key and not its value
	_, v4 := User["Name"]

	fmt.Println(k1, v1)
	fmt.Println(k2, v2)
	fmt.Println(k3, v3)
	fmt.Println(v4)

}
