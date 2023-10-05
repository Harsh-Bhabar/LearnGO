package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)

	languages["A"] = "first"
	languages["B"] = "second"
	languages["C"] = "third"
	languages["D"] = "forth"

	fmt.Println(languages)
	fmt.Println(languages["A"]) // gives its value
	fmt.Println(languages["Z"]) // gives empty

	// delete value

	delete(languages, "A")
	fmt.Println(languages)

	// looping

	for key, value := range languages {
		fmt.Printf("key is %v and value is %v\n", key, value)
	}
}
