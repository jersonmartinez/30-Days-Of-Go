package main

import "fmt"

func vowelsCounter(sentence string) (int, int, int, int, int) {
	counter_a := 0
	counter_e := 0
	counter_i := 0
	counter_o := 0
	counter_u := 0
	for _, value := range sentence {
		variable := string(value)
		switch variable {
		case "a":
			counter_a++
		case "e":
			counter_e++
		case "i":
			counter_i++
		case "o":
			counter_o++
		case "u":
			counter_u++
		}
	}
	return counter_a, counter_e, counter_i, counter_o, counter_u
}

func main() {

	sentence := "For example, this is a litter sentence with a lot of vowels"
	a, e, i, o, u := vowelsCounter(sentence)
	fmt.Printf("Phrase '%s' has: \n", sentence)
	fmt.Printf("%d vowel a \n", a)
	fmt.Printf("%d vowel e \n", e)
	fmt.Printf("%d vowel i \n", i)
	fmt.Printf("%d vowel o \n", o)
	fmt.Printf("%d vowel u \n", u)
}
