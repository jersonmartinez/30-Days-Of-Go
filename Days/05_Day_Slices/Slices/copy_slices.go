package main

import "fmt"

func main() {
	OS := []string{"Ubuntu", "Debian", "Windows 10", "Windows 11", "Fedora", "Mac OS X"}
	var OS_Copy = make([]string, len(OS))

	fmt.Println(OS)

	copy(OS_Copy, OS)

	OS = append(OS[:1], OS[2:]...)

	fmt.Println(OS_Copy)
	fmt.Println(OS)
}
