package main

import "fmt"

func main() {
	/* Palabra reservada + Nombre de variable + Tipo de dato + Valor */
	const PI float64 = 3.14
	const PI2 = 3.14

	fmt.Println("Pi:", PI, "\nPi2: ", PI2)

	/* Constante con nombre de variable y valor */
	const (
		a = 1
		b = 2
	)

	fmt.Println("a:", a, "\nb:", b)
}
