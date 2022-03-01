package main

import "fmt"

func main() {
	/* Palabra reservada + Nombre de variable + Tipo de dato + Valor */
	var A int = 10
	var B int = 20

	Suma := A + B
	Resta := A - B
	Multiplicacion := A * B
	Division := A / B
	Modulo := A % B

	A++
	B--

	fmt.Println("Suma:", Suma)
	fmt.Println("Resta:", Resta)
	fmt.Println("Multiplicacion:", Multiplicacion)
	fmt.Println("Division:", Division)
	fmt.Println("Modulo:", Modulo)
	fmt.Println("Incremental:", A)
	fmt.Println("Decremental:", B)
}
