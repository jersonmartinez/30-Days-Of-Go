package main

import "fmt"

func main() {
	/* Declara la variable "Base", identifica el tipo de dato respecto al valor */
	Base := 10          //Primera forma
	var Altura int = 20 //Segunda forma
	var Area int        //Go no compila si las variables no son usadas

	fmt.Println("Base: ", Base, " | Altura: ", Altura, " | Area: ", Area)

	/*Si no hay valores definidos, Go le asigna un vacío.*/
	var a int     //0
	var b float64 //0.0
	var c string  //""
	var d bool    //false

	fmt.Println("a: ", a, " | b: ", b, " | c: ", c, " | d: ", d)
}
