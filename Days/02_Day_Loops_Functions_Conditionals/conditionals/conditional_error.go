package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	value, err := strconv.Atoi("100")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", value)
		log.Fatal(value)
	}
}
