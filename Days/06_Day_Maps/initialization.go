package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{
		"Username": "jersonmartinez",
		"Name":     "Jerson Antonio",
		"Lastname": "Martínez Moreno",
	}

	b := map[string]int{
		"Day":          6,
		"Month":        4,
		"Year":         1996,
		"Magic number": 9,
	}

	fmt.Printf("a = \t%v\n", a)
	fmt.Printf("b = \t%v\n", b)

}
