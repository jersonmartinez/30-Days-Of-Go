package main

import (
	"fmt"
)

// main function
func main() {

	/*Declarando un slice bidimensional*/
	slice_of_slices := make([][]int, 3)

	for i := 0; i < 3; i++ {
		/*Declarando un slice dentro de cada elemento del slice padre*/
		slice_of_slices[i] = make([]int, 3)

		/*Asignando valores a cada elemento del slice*/
		for j := 0; j < 3; j++ {
			slice_of_slices[i][j] = i * j
		}
	}

	/*Imprimiendo los slices hijos*/
	for i := 0; i < 3; i++ {
		fmt.Println(slice_of_slices[i])
	}

	/*Imprimiendo todo el slice*/
	fmt.Println("Slice of slices: ", slice_of_slices)
}
