package main

import (
	"fmt"
	"math"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * a
}

func dobleReturnValue(a int) (b, c int) {
	return a, a * a
}

func areaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}
func areaRectangulo(base float64, altura float64) float64 {
	return base * altura
}

func areaTrapezoide(B float64, b float64, h float64) float64 {
	return h * (B + b) / 2
}

func main() {

	normalFunction("Bueeenas")
	tripleArgument(1, 4, "leo")
	fmt.Println(returnValue(7))
	fmt.Println(dobleReturnValue(4))
	fmt.Printf("Circulo %.2f \n", areaCirculo(2))
	fmt.Printf("Rectangulo %.2f \n", areaRectangulo(5, 10))
	fmt.Printf("Trapezoide %.2f \n", areaTrapezoide(10, 5, 3))

}
