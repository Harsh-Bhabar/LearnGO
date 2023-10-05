package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var names = []string{}
	fmt.Printf("Type of names is - %T\n", names)

	names = append(names, "a", "b", "c", "d", "e", "f", "g", "h")
	fmt.Println(names)

	names = append(names[1:]) // from 1 to end
	fmt.Println(names)

	names = append(names[1:5]) // from 1 to n-1
	fmt.Println(names)

	names = append(names[:2]) // from start to n
	fmt.Println(names)

	highscores := make([]int, 4)
	highscores[0] = 100
	highscores[1] = 200
	highscores[2] = 300
	highscores[3] = 400
	// highscores[4] = 500                            // will throw out of bound
	highscores = append(highscores, 1, 2, 3, 4, 5) // will rellocate the memory and size
	fmt.Println(highscores, "len", len(highscores))

	fmt.Println(sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	// delete from slices
	courses := []string{}
	courses = append(courses, "react", "js", "java", "html")
	fmt.Println(courses)

	index := 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
