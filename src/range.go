package main

import "fmt"

func isPalindrom(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es palidromo")
	} else {
		fmt.Println("No es palidromo")
	}
}

func main() {

	slice := []string{"hola", "que", "tal"}

	for i := range slice {
		fmt.Println(i)
	}

	isPalindrom("ama")

}
