package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {

	var wg sync.WaitGroup

	fmt.Println("Hello")

	wg.Add(1)

	go say("World", &wg)

	wg.Wait()

	wg.Add(1)

	go func(text string) {
		defer wg.Done()
		fmt.Println(text)
	}("Adios")

	wg.Wait()

	// time.Sleep(time.Second * 1)

	// NOTA: los channels son mejor

}
