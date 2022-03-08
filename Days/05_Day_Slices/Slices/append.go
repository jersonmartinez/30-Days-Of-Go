package main

import "fmt"

func main() {
	var marcasDeCoches = make([]string, 2)
	marcasDeCoches[0] = "Aprendiendo"
	marcasDeCoches[1] = "Go"

	marcasDeCoches = append(marcasDeCoches, "en", "la", "nube")
	fmt.Println(marcasDeCoches)
}
