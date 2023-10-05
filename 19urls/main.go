package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:3000/learn?coursename=react&paymentId=124124aefaf"

func main() {
	fmt.Println("URLS")
	fmt.Println(myUrl)

	result, err := url.Parse(myUrl)
	errorcheck(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// query params

	qparams := result.Query()
	fmt.Printf("Qparams type is %T\n", qparams)
	fmt.Println(qparams)

	for _, val := range qparams {
		fmt.Println("VAL - ", val)
	}

	//constuct URLs

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco-dev",
		Path:    "/learn",
		RawPath: "url=harsh",
	}

	finalUrl := partsOfUrl.String()
	fmt.Println(finalUrl)
}

func errorcheck(err error) {
	if err != nil {
		panic(err)
	}
}
