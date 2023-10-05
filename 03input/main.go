package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome!"
	fmt.Println(welcome)

	// takng input

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name - ")

	// comma-ok syntax / try-catch / err ok

	name, _ := reader.ReadString('\n')
	fmt.Println("Hey,", name) // anything will be treated as string

}
