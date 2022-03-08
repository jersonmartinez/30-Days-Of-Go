package main

import "fmt"

func main() {

	mySlice := []int{}

	for i := 0; i < 100; i += 1 {
		mySlice = append(mySlice, i+1)
	}

	fmt.Printf("%+v", mySlice)
}
