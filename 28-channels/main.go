package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in go lang")

	myCh := make(chan int, 2) //now adding 2 buffer so we can listen 2 values at a time
	wg := &sync.WaitGroup{}

	//store 5 into channel
	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2) //adding two go routine

	// <-chan mens receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChannelOpen := <-myCh

		fmt.Println("Is Channel Open:", isChannelOpen)
		fmt.Println("Value:", val)
		wg.Done()
	}(myCh, wg)

	// chan<- mens send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
