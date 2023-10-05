package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Current time is = ")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 Monday"))
}
