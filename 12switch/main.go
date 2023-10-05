package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch")

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1

	fmt.Println(diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Dice value is 1, you can Open.")
	case 2:
		fmt.Println("Dice value is 2, you can move 2 spot.")
	case 3:
		fmt.Println("Dice value is 3, you can move 3 spot.")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4, you can move 4 spot.")
	case 5:
		fmt.Println("Dice value is 5, you can move 5 spot.")
	case 6:
		fmt.Println("Dice value is 6, you can move 6 spot.")
	default:
		fmt.Println("WTF Bro")
	}

}
