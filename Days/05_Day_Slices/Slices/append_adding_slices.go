// Golang program to demonstrate the
// example of append() function

package main

import (
	"fmt"
)

func main() {
	// Creating int and string slices
	s1 := []int{10, 20, 30}
	s2 := []string{"Hello", "World"}

	// Printing types and values of slices
	fmt.Printf("%T, %v\n", s1, s1)
	fmt.Printf("%T, %q\n", s2, s2)

	// Appending the elements
	s1 = append(s1, 40, 50)
	s2 = append(s2, "How are you?", "Boys", "and", "Girls")

	fmt.Println("After appending...")
	fmt.Printf("%T, %v\n", s1, s1)
	fmt.Printf("%T, %q\n", s2, s2)
}
