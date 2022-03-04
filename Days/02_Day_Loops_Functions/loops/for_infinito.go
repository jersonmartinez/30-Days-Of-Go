package main

import "fmt"

func main() {
	counter := 0
	for {
		fmt.Println("Forever: ", counter)
		counter++
	}
}
