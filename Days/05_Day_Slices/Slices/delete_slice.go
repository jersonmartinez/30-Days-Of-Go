package main

import "fmt"

func main() {
	OS := []string{"Ubuntu", "Debian", "Windows 10", "Windows 11", "Fedora", "Mac OS X"}
	fmt.Println(OS)

	indice := 2 // Windows 10

	OS = append(OS[:indice], OS[indice+1:]...)
	fmt.Println(OS)
}
