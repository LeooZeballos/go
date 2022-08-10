package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram *= 2
}



func main() {

	a := 50
	b := &a

	fmt.Println(a, b, *b)

	*b = 100
	fmt.Println(a, *b)

	myPC := pc{ram: 16, disk: 500, brand: "yes"}
	fmt.Println(myPC)
	myPC.ping()
	myPC.duplicateRAM()
	myPC.duplicateRAM()

	fmt.Println(myPC)

}
