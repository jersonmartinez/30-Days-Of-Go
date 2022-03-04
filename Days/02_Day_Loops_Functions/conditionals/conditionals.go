package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 es par")
	} else {
		fmt.Println("7 es impar")
	}

	if 8%4 == 0 {
		fmt.Println("8 es divisible de 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "es negativa")
	} else if num < 10 {
		fmt.Println(num, "tiene 1 dígito")
	} else {
		fmt.Println(num, "tiene multiples dígitos")
	}
}
