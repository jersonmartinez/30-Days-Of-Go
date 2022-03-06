package main

import "fmt"

func main() {
	countries := [...]string{"India", "Canada", "Japan", "Germany", "Italy"}

	fmt.Printf("Countries: %v\n", countries)

	fmt.Printf(":2 %v\n", countries[:2])

	fmt.Printf("1:3 %v\n", countries[1:3])

	fmt.Printf("2: %v\n", countries[2:])

	fmt.Printf("2:5 %v\n", countries[2:5])

	fmt.Printf("0:3 %v\n", countries[0:3])

	fmt.Printf("Last element: %v\n", countries[len(countries)-1])

	fmt.Printf("All elements: %v\n", countries[0:len(countries)])
	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:len(countries)])

	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])
}
