package main

import "fmt"

func main() {
	fmt.Println("If else")

	age := 18
	var result string

	if age < 18 {
		result = "Minor"
	} else if age == 18 {
		result = "PERFECT"
	} else {
		result = "Not minor"
	}

	fmt.Println(result)

	// assigning and using at same time

	if time := 100; time < 10 {
		fmt.Println("Time is less than 10")
	} else {
		fmt.Println("Time is greater than 10")
	}
}
