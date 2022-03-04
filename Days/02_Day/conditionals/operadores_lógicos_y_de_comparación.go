package main

import "fmt"

func main() {

	// Operadores lógicos y de comparación
	// Son operadores que nos permiten hacer una comparación de condiciones y en caso de cumplirse como sino ejecutarán un código determinado. Si se cumple es VERDADERO/TRUE y si no se cumple son FALSO/FALSE.

	// Empecemos con los operadores de comparación:

	// Operadores de comparación
	// Son aquellos que retornan TRUE o FALSE en caso de cumplirse o no una expresión. Son los siguientes:

	//Retorna TRUE si valor1 y valor2 son exactamente iguales.
	valor1 := 10
	valor2 := 10
	fmt.Println(valor1 == valor2)

	//Retorna TRUE si valor1 es diferente de valor2.
	valor1 = 5
	fmt.Println(valor1 != valor2)

	//Retorna TRUE si valor1 es menor que valor2
	fmt.Println(valor1 < valor2)

	//Retorna TRUE si valor1 es mayor que valor2
	valor1 = 15
	fmt.Println(valor1 > valor2)

	//Retorna TRUE si valor1 es igual o mayor que valor2
	fmt.Println(valor1 >= valor2)

	//Retorna TRUE si valor1 es menor o igual que valor2.
	valor2 = 20
	fmt.Println(valor1 <= valor2)

	// Operadores lógicos
	// Son aquellos que retorna TRUE o FALSE
	// si cumplen o no una condición utilizando puertas lógicas.

	// Operador AND:
	// Este operador indica que todas las condiciones declaradas
	// deben cumplirse para poderse marcar como TRUE.
	// En Go, se utiliza este símbolo &&.

	// Esto retornará TRUE porque tanto la primera
	// como la segunda condición son verdaderas.
	fmt.Println(1 > 0 && 2 > 0)

	// Esto retornará FALSE porque una de las condiciones
	// no es verdadera.
	fmt.Println(2 < 0 && 1 > 0)

	// Operador OR:
	// Este operador indica que al menos una de las condiciones
	// debe cumplirse para marcarse como TRUE.
	// En Go, se representa con el símbolo ||.

	// Esto retornará TRUE porque la segunda condición se cumple,
	// a pesar que la primera no.
	fmt.Println(2 < 0 || 1 > 0)

	//Operador NOT:
	// Este operador retornará el opuesto al boleano
	// que está dentro de la variable

	myBool := true
	// Esto retornará false
	fmt.Println(!myBool)
}
