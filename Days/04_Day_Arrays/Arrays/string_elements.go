package main

import "fmt"

func main() {
	/*First array*/
	var array1 [2]string
	array1[0] = "Hello"
	array1[1] = "World"

	fmt.Println(array1[0], array1[1])
	fmt.Println(array1)

	/*Second array*/
	array2 := [2]string{"Hello", "World"}

	fmt.Println("\n", array2[0], array2[1])
	fmt.Println(array2)
}
