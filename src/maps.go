package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 16

	fmt.Println(m)

	// recorrer Map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// encontrar valor -> default (si no lo encuentra) = zero, ok = true o false (si lo encontro o no)
	value, ok := m["Jose"]
	fmt.Println(value, ok)

}
