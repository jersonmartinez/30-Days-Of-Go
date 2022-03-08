package main

import "fmt"

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	OS := []string{"Ubuntu", "Debian", "Windows 10", "Windows 11", "Fedora", "Mac OS X"}
	fmt.Println(OS)

	OS = remove(OS, 2)

	fmt.Println(OS)
}
