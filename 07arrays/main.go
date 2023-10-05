package main

import "fmt"

func main() {
	fmt.Println("arrays")

	var names [5]string
	names[0] = "harsh"
	names[1] = "harsh1"
	names[2] = "harsh2"
	names[4] = "hars2"

	fmt.Println(names)      // a b c -spave- d
	fmt.Println(len(names)) // gives 5 -> the actual size

}
