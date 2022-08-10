package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Ambos son verdad")
	}

	// With OR
	if valor1 == 1 || valor2 == 2 {
		fmt.Println("Uno es verdad")
	}

	// Texto a número
	value, err := strconv.Atoi("54")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)

}
