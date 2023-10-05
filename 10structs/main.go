package main

import "fmt"

type User struct {
	Name       string
	Email      string
	Autorizedd bool
	Age        int
}

func main() {
	fmt.Println("Structs")

	// NO INHERITENCE or such components IN GO

	harsh := User{"Harsh", "h@gmail.com", true, 24}
	fmt.Println(harsh)

	fmt.Printf("hars details - %+v\n", harsh)
	fmt.Printf("Name is - %v\n", harsh.Name)

}
