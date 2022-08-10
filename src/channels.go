package main

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}

func main() {

	// Channels? Organizar las goroutines

	// Especificar entrada o salida

	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)

	fmt.Println(<-c)

}
