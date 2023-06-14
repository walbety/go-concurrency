package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	waitGroupChallenge()
}

func waitGroupExample() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"gama",
		"delta",
		"pi",
		"zeta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	fmt.Println("done!")
}

var msg string = "Hello, world!"

func waitGroupChallenge() {

	var msgs = []string{
		"Hello, world!",
		"Hello, universe!",
		"Hello, cosmos!",
	}

	var wg sync.WaitGroup
	wg.Add(len(msgs))

	for _, msg := range msgs {
		go func(s string) {
			updateMessage(s)
			printMessage()
			defer wg.Done()
		}(msg)
	}
	wg.Wait()

	fmt.Println("challenge done!")
}

func printMessage() {
	fmt.Println(msg)
}

func updateMessage(s string) {
	msg = s
}
