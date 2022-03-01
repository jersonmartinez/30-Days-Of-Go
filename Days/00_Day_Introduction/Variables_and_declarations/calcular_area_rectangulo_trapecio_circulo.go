package main

import (
	"fmt"
	"math"
)

func main() {
	/* Declaracion de variables base */
	Base := 10
	BaseMayor := 20
	Altura := 30
	Radio := 20

	fmt.Println("El área del rectángulo es:", Base*Altura)
	fmt.Println("El área del trapecio es:", ((BaseMayor+Base)*Altura)/2)
	fmt.Println("El área del círculo es:", math.Pi*math.Pow(float64(Radio), 2))
}
