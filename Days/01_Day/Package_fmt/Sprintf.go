package main

import "fmt"

func main() {

	Fistname := "Jerson"
	Lastname := "Martínez"

	/*
		Fuente: https://pkg.go.dev/fmt#Sprintf
		Sprintf aplica un formato a una cadena de caracteres y devuelve una cadena de caracteres.
	*/
	Message := fmt.Sprintf("Hola, %s %s", Fistname, Lastname)

	fmt.Println(Message)
}
