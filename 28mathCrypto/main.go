package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Maths")

	var first int = 1
	var second float64 = 4

	fmt.Println("sum", first+int(second))

	// random number

	// math/rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1) // math.rand -> (0 to 5) + 1

	// crypto/rand
	myrand, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myrand)
}
