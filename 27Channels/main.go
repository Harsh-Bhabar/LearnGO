package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels")

	myChannel := make(chan int, 1)
	wg := &sync.WaitGroup{}

	// myChannel <- 5
	// fmt.Println(<-myChannel)

	wg.Add(2)

	// means value going outside the box/channel
	// RECEIVE ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChannel
		fmt.Println(val, isChannelOpen)
		// fmt.Println(<-ch)
		wg.Done()
	}(myChannel, wg)

	// means value going into the box/channel
	//SEND ONLY
	go func(ch chan int, wg *sync.WaitGroup) {
		// ch <- 5
		ch <- 6
		close(myChannel)
		wg.Done()
	}(myChannel, wg)

	wg.Wait()

}
