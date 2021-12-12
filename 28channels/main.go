package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("This is all about channels in Golang")

	myChannel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myChannel <- 5
	// fmt.Println(<-myChannel)
	wg.Add(2)

	// recieve only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(myChannel)
		val, isChannelOpen := <-myChannel

		fmt.Println(val)
		fmt.Println(isChannelOpen)

		// fmt.Println(<-myChannel)

		wg.Done()
	}(myChannel, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChannel <- 0
		close(myChannel)
		// myChannel <- 0
		// myChannel <- 6
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
