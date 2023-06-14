package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {

	for {
		// print a got data message
		i := <-ch
		fmt.Println("Got", i, "from channel")

		// simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}

func main() {

	/*
		Buffered channels are useful when you know how many go routines you've launch.
		Or we want to limit the number of go routines we launch.
		Or we want to limit the amount of work that's queued up, which is our case bellow.
	*/

	// ch := make(chan int) // unbuffered channel
	ch := make(chan int, 10) // buffered channel

	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("sending", i, "to channel....")

		ch <- i

		fmt.Println("sent", i, "to channel!")
	}

	fmt.Println("Done!")

	close(ch)

}
