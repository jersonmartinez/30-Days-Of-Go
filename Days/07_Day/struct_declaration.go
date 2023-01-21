package main

import "fmt"

type Point struct {
	x float64
	y float64
}

func main() {
	p := Point{x: 3.0, y: 4.0}

	fmt.Println("X:", p.x)
	fmt.Println("Y:", p.y)
}
