package main

import "fmt"

func main() {
	x := 10
	y := 51

	fmt.Println("Suma:", x+y)
	fmt.Println("Resta:", y-x)
	fmt.Println("Multiplicación:", x*y)
	fmt.Println("División:", y/x)
	fmt.Println("Módulo:", y%x)
	x++
	fmt.Println("Incremental:", x)
	y--
	fmt.Println("Decremental:", y)

	// Rectángulo
	baseRectangulo := 20
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("El Area del Rectángulo es :", areaRectangulo)

	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10

	areaCirculo := PI * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

	// Trapecio
	var baseUno float64 = 6
	var baseDos float64 = 15
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Area del Trapecio es :", areaTrapecio)

}
