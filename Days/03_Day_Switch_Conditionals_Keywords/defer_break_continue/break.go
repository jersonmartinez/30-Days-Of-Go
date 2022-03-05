package main

import "fmt"

func main() {
	var i uint8 = 0
	for i < 10 {
		fmt.Println(i)
		i++

		if i == 8 {
			fmt.Println("Break!")
			break
		}
	}
}
