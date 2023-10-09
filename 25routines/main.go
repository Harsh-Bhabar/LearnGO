package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var waitGroup sync.WaitGroup // pointers
var mutex sync.Mutex         //pointer

var signals = []string{"Test"}

func main() {
	fmt.Println("Routines")
	// go greeter("world")
	// greeter("world2")

	websites := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.fb.com",
		"https://www.github.com",
	}

	for _, site := range websites {
		go getStatusCode(site)
		waitGroup.Add(1)
	}

	waitGroup.Wait()

	fmt.Println(signals)

}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(30 * time.Millisecond)
// 		fmt.Printf("Hello %v\n", s)
// 	}
// }

func getStatusCode(endpoint string) {

	defer waitGroup.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Status Code for %v is %v\n", endpoint, res.StatusCode)
		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()
	}

}
