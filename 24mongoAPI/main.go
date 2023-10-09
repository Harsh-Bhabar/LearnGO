package main

import (
	"fmt"
	"mongoApis/router"
	"net/http"
)

func main() {
	fmt.Println("Final")

	router := router.Router()

	fmt.Println("Staring server...")
	http.ListenAndServe(":8000", router)

	fmt.Println("Listening...")
}
