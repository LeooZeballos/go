package main

import "fmt"

func main() {

	helloMessage := "Hello World!"

	// Println
	fmt.Println(helloMessage)

	// Printf
	nombre := "Leonel"
	fmt.Printf("%s tiene 21 años\n", nombre)

	// Sprintf
	message := fmt.Sprintf("%s tiene mucha plata, dea\n", nombre)
	fmt.Print(message)

}
