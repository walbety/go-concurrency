package main

import (
	"fmt"
	"strings"
)

func shout(ping chan string, pong chan string) {
	for {
		s := <-ping

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Printf("Type something and press ENTER (enter Q to quit)")

	for {
		fmt.Print("-> ")

		var userInput string

		_, _ = fmt.Scanln(&userInput)

		if userInput == "q" {
			break
		}

		ping <- userInput

		response := <-pong

		fmt.Println(response)

	}

	fmt.Println("All done. Closing channels.")

	close(ping)
	close(pong)

}
