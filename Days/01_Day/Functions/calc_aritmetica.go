package main

import "fmt"

func suma(A, B int) int           { return A + B }
func resta(A, B int) int          { return A - B }
func multiplicacion(A, B int) int { return A * B }
func division(A, B int) int       { return A / B }

func main() {
	var A int = 20
	var B int = 10

	Suma := suma(A, B)
	Resta := resta(A, B)
	Multiplicacion := multiplicacion(A, B)
	Division := division(A, B)

	fmt.Println("Suma:", Suma)
	fmt.Println("Resta:", Resta)
	fmt.Println("Multiplicación:", Multiplicacion)
	fmt.Println("División:", Division)
}
