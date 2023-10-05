package main

import "fmt"

const LoginTime = "12334151" // first letter as capital makes it PUBLIC

func main() {

	//variables

	var username string = "Harsh"
	fmt.Println(username)
	fmt.Printf("Variable is of type - %T\n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type - %T\n", isLoggedin)

	var unsignedSmallInt uint = 255
	fmt.Println(unsignedSmallInt)
	fmt.Printf("Variable is of type - %T\n", unsignedSmallInt)

	var normalInt = 2412414
	fmt.Println(normalInt)
	fmt.Printf("Variable is of type - %T\n", normalInt)

	var smallFloat float32 = 255.321241241241241441
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type - %T\n", smallFloat)

	var longFloat float64 = 255.321241241241241441
	fmt.Println(longFloat)
	fmt.Printf("Variable is of type - %T\n", longFloat)

	//default values and aliases

	var anotherInt int
	fmt.Println(anotherInt)
	fmt.Printf("Variable is of type - %T\n", anotherInt)

	var anotherStr string
	fmt.Println(anotherStr)
	fmt.Printf("Variable is of type - %T\n", anotherStr)

	// implicit type
	var website = "somewebsite"
	fmt.Println(website)
	fmt.Printf("variable is of tyoe - %T\n", website) // you wont be able to change it to another type

	// walrus operator to declare vars
	numberOfUsers := 2000
	fmt.Println(numberOfUsers)
	numberOfUsers = 5000
	fmt.Println(numberOfUsers)

	//accessing public vars
	fmt.Println(LoginTime)

}
