package main

import "fmt"

func main() {
	if isPair(6) {
		fmt.Println("Número par")
	} else {
		fmt.Println("Número impar")
	}
	if isValidUser("Antonio5", "Password") {
		fmt.Println("Credentiales válidas")
	} else {
		fmt.Println("Credentiales inválidas")
	}
}

func isPair(num int) bool {
	return num%2 == 0
}

func isValidUser(userName, pass string) bool {
	return userName == "Antonio" && pass == "Password"
}
