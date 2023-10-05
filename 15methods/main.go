package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) GetAge() int {
	return u.Age
}

func (u User) setNewName() {
	u.Name = "Mrinal"
	fmt.Println(u)
}

func main() {
	fmt.Println("MEthods")

	harsh := User{"harsh", "hb@gmail.com", 24}
	age := harsh.GetAge()
	fmt.Printf("Age is %v\n", age)

	// passes a copy of the obj, not actual obj
	harsh.setNewName()
	fmt.Println(harsh)

}
