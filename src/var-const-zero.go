package main

import "fmt"

func main() {

	// Declaración de constantes
	const PI float64 = 3.14159
	const PI2 = 3.14

	fmt.Println("PI", PI)
	fmt.Println("PI2", PI2)

	// Declaración de variables enteras
	base := 12
	var altura int = 15
	var area int

	fmt.Println(base, altura, area)

	// Zero values (por defecto)
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area del cuadrado", areaCuadrado)

}
