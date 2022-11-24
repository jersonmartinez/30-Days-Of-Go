package main

import (
	"fmt"
)

func main() {

	var a = make(map[string]string)

	a["Username"] = "jersonmartinez"
	a["Name"] = "Jerson Antonio"
	a["Lastname"] = "Martínez Moreno"

	b := make(map[string]int)

	b["Day"] = 6
	b["Month"] = 4
	b["Year"] = 1996
	b["Magic number"] = 9

	fmt.Printf("a = \t%v\n", a)
	fmt.Printf("b = \t%v\n", b)

}
