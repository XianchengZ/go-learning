package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	
	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().
	var waitGroup sync.WaitGroup
	
	msg = "Hello, world!"
	
	waitGroup.Add(1)
	go updateMessage("Hello, universe!", &waitGroup)
	waitGroup.Wait()
	printMessage()
	
	waitGroup.Add(1)
	go updateMessage("Hello, cosmos!", &waitGroup)
	waitGroup.Wait()
	printMessage()
	
	waitGroup.Add(1)
	go updateMessage("Hello, world!", &waitGroup)
	waitGroup.Wait()
	printMessage()

}
