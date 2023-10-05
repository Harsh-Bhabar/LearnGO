package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("HTTP")

	response, err := http.Get(url)
	errorCheck(err)

	fmt.Println("response body - ", response)
	fmt.Printf("response type  %T-\n", response)

	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)
	errorCheck(err)

	content := string(dataBytes)
	fmt.Println(content)
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
