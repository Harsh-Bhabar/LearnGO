package main

import "fmt"

func main() {
	defer fmt.Println("LOL")
	fmt.Println("defer")

	defer fmt.Println("second")
	defer fmt.Println("third")
	mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
