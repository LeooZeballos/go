package main

import "fmt"

func main() {

	a := 4
	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("No")
	}

	// Sin condicion
	value := 200
	switch {
	case value > 100:
		fmt.Println("Mayor a 100")
	case value < 0:
		fmt.Println("Menor a 0")
	default:
		fmt.Println("Default")
	}

}
