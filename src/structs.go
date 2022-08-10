package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {

	myCar := car{brand: "Tesla", year: 2022}
	otherCar := car{brand: "Ford"}
	fmt.Println(myCar)
	fmt.Println(otherCar)

}
