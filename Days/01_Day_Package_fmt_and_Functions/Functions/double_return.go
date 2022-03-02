package main

import "fmt"

func cuadrado(n int) (x, y int) {
	return n, n * n
}

func main() {
	var N int = 20

	Value, Cuadrado := cuadrado(N)
	fmt.Printf("Value: %d, Cuadrado %d\n", Value, Cuadrado)

	/*Solo obtiene un valor*/
	_, Value = cuadrado(N)
	fmt.Printf("Cuadrado: %d\n", Value)
}
