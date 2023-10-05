package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	myCourses := []course{
		{"React", 299, "harsh123", []string{"a", "b", "c"}},
		{"HTML", 399, "harsh123", []string{"a", "b", "c"}},
		{"JS", 499, "harsh123", nil},
	}

	// package as JSON data

	// finalJson, err := json.Marshal(myCourses)
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")
	errHandle(err)
	fmt.Println(string(finalJson))
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
                "courseName": "HTML",
                "Price": 399,
                "tags": [
                        "a",
                        "b",
                        "c"
                ]
        }
	`)

	var myCourses course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &myCourses)
		fmt.Printf("DATA is - %#v\n", myCourses)
	} else {
		fmt.Println("Json is invalid")
	}

	// when we just want to add data to key value pair

	var mydata map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &mydata)
	fmt.Printf("DATA is - %#v\n", mydata)

	for k, v := range mydata {
		fmt.Printf("key - %v and val - %v\n", k, v)
	}

}

func errHandle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("JSON")
	// EncodeJson()
	DecodeJson()
}
