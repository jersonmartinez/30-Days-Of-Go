package main

import "fmt"

func main() {
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	for i := 0; i < 10; i++ {

		if i == 2 {
			fmt.Println("es 2")
			continue
		}

		fmt.Println(i)

		if i == 5 {
			break
		}
	}
}
