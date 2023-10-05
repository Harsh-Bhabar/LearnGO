package main

import "fmt"

func main() {
	fmt.Println("Functions")
	hello()

	sum := add(2, 5)
	fmt.Printf("sum is %v\n", sum)

	sum2, sum2Str := proadd(1, 2, 2, 4, 4, 5, 6, 2, 5, 5, 5)
	fmt.Printf("sum is %v and returned str is %v\n", sum2, sum2Str)

}

func proadd(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "PROADD"
}

func add(val1 int, val2 int) int {
	return val1 + val2
}

func hello() {
	fmt.Println("hello")
}
