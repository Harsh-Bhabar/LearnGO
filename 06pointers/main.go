package main

import "fmt"

func main() {
	fmt.Println("Poninter")

	var ptr *int
	fmt.Println("value of ptr is - ", ptr) // nil

	myNum := 23

	var myNumPtr = &myNum                   // providing the reference of the myNum
	fmt.Println(myNum, *myNumPtr, myNumPtr) // 23 23 location

	*myNumPtr = *myNumPtr + 2 // operation will be performed on the actual values
	fmt.Println(myNum)

}
