package main

import "fmt"

func multiplication(a, b int) int {
	return (a * b)
}

func show() {
	fmt.Println("Hello, my name is Antonio!")
}

func main() {

	fmt.Println("Result is:", multiplication(10, 10))

	defer fmt.Println("Result is:", multiplication(15, 15))

	show()
}
