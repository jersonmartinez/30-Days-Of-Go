package main

import "fmt"

func main() {

	Firstname := "Jerson"
	Lastname := "Martínez"

	Year := 2030

	/*
		Imprime en pantalla el mensaje y los valores de las variables
		%s: String
		%d: Entero
		%f: Flotante
		%t: Booleano

		%v: Valor de la variable
		%T: Tipo de dato de la variable

		%p: Dirección de memoria de la variable

		%b: Binario
		%c: Caracter
		%o: Octal
		%x: Hexadecimal

		%q: Cadena de caracteres
		%e: Notación científica
		%g: Notación científica
		%f: Notación científica

		%+v: Valor de la variable con su tipo de dato
		%#v: Valor de la variable con su tipo de dato
	*/

	fmt.Printf("Hola, %s %s\n", Firstname, Lastname)
	fmt.Printf("Vengo desde el %d\n", Year)

	fmt.Printf("Dirección de memorias de las variables: Firstname %p, Lastname %p, Year %p\n", &Firstname, &Lastname, &Year)

}
