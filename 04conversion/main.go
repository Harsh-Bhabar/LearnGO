package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your age - ")

	ageInput, _ := reader.ReadString('\n') // here ageInput == "24\n"
	fmt.Println("Age is -", ageInput)

	age, err := strconv.ParseFloat(strings.TrimSpace(ageInput), 64)

	if err != nil {
		fmt.Println("err - ", err)
	} else {
		fmt.Println("Adding 1 to yout age - ", age+1)
	}

}
