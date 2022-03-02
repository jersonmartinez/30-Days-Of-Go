package main

import (
	"fmt"
	"math"
)

func áreaCírculo(radio float64) float64 {
	return math.Pi * radio * radio
}
func áreaRectángulo(base, altura float64) float64 {
	return base * altura
}

func áreaTrapecio(B, b, h float64) float64 {
	return h * (B + b) / 2
}

func main() {
	fmt.Printf("Círculo: %.2f \n", áreaCírculo(2))
	fmt.Printf("Rectángulo: %.2f \n", áreaRectángulo(5, 10))
	fmt.Printf("Trapecio: %.2f \n", áreaTrapecio(10, 5, 3))
}
