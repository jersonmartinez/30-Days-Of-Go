package main

import (
	"fmt"
)

func main() {
	var array1 = [...]int{1, 2, 3}
	array2 := [...]int{4, 5, 6, 7, 8}
	array_string := [...]string{"Hello", "World"}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array_string)
}
