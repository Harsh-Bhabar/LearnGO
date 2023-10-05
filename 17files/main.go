package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files")

	content := "This needs to go in file."

	file, err := os.Create("./myfile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is - ", length)
	defer file.Close()

	readFile("./myfile.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Data in file", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
