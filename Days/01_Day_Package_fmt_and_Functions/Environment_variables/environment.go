package main

import (
	"fmt"
	"os"
)

func main() {

	var (
		home   = os.Getenv("HOME")
		user   = os.Getenv("USER")
		gopath = os.Getenv("GOPATH")
	)

	fmt.Println("HOME:", home)
	fmt.Println("USER:", user)
	fmt.Println("GOPATH:", gopath)
}
