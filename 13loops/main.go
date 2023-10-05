package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{}

	days = append(days, "mon", "tue", "wed", "thu", "fri")
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for i, day := range days {
		fmt.Printf("index and day is %v and %v\n", i, day)
	}

	// while

	val := 1
	for val < 10 {

		// break continue goto

		// if val == 5 {
		// 	break
		// }
		if val == 6 {
			val++
			continue
		}
		if val == 7 {
			goto gotoStatement
		}

		fmt.Println(val)
		val++
	}

gotoStatement:
	fmt.Println("Came to to goto")
}
