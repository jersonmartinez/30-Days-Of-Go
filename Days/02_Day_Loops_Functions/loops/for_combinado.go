package main

import "fmt"

func main() {
	/*
		Imprime los números del 1 al 3
	*/
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	/*
		Imprime los números del 7 al 9
	*/
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	/*
		Realiza un ciclo infinito, sin embargo, lo detiene a la primera.
	*/
	for {
		fmt.Println("loop")
		break
	}

	/*
		Imprime solo números impares, ya que omite los números pares por medio del módulo.
	*/
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
