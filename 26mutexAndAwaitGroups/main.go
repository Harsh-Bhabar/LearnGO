package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func main() {
	fmt.Println("Race condition")

	waitGroup := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	scores := []int{0}

	waitGroup.Add(3)
	// waitGroup.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Rountine-1")
		mut.Lock()
		scores = append(scores, 1)
		mut.Unlock()
		wg.Done()
	}(waitGroup, mutex)

	// waitGroup.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Rountine-2")
		mut.Lock()
		scores = append(scores, 2)
		mut.Unlock()
		wg.Done()
	}(waitGroup, mutex)

	// waitGroup.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Rountine-3")
		mut.Lock()
		scores = append(scores, 3)
		mut.Unlock()
		wg.Done()
	}(waitGroup, mutex)

	waitGroup.Wait()

	fmt.Println(scores)

}
