package main

import (
	"fmt"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("First")
	case 2:
		fmt.Println("Second")
	case 3:
		fmt.Println("Third")
	}
}
