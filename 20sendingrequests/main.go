package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Requests")
	// makeGetRequest()
	// makeJsonPostRequest()
	performPostFormRequest()
}

func makeGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	errHandle(err)

	defer response.Body.Close()

	fmt.Println("Status code - ", response.StatusCode)
	fmt.Println("Content Len - ", response.ContentLength)

	// method 1
	// content, err := ioutil.ReadAll(response.Body)
	// errHandle(err)
	// fmt.Println(string(content))

	// method 2
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	responseString.Write(content)
	fmt.Println(responseString.String())

}

func makeJsonPostRequest() {
	const myUrl = "http://localhost:8000/post"
	//fake json
	requestBody := strings.NewReader(`
		{
			"name" : "harsh",
			"age": 24
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	errHandle(err)

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	errHandle(err)
	fmt.Println(string(content))

}

func performPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstName", "Harsh")
	data.Add("lastName", "Bhabar")
	data.Add("age", "24")

	response, err := http.PostForm(myUrl, data)
	errHandle(err)

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	errHandle(err)

	fmt.Println(string(content))

}

func errHandle(err error) {
	if err != nil {
		panic(err)
	}
}
