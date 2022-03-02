package main

import (
	"log"
	"os"
)

func main() {

	key := "NEW_VAR"
	value := "This is the first value in this environment value!"

	os.Setenv(key, value)

	val := GetEnvDefault(key, value)

	log.Println("\nEl valor es: " + val)

	os.Unsetenv(key)

	val = GetEnvDefault(key, value)

	log.Println("\nEl valor por default es: " + val)

}

// GetEnvDefault buscar las vars
func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
