package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println(s)
}

func main() {

	var waitGroup sync.WaitGroup
	
	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	waitGroup.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &waitGroup)
	}

	waitGroup.Wait() // wait until the wait group down to 0


	waitGroup.Add(1)
	printSomething("This is the second thing to be printed", &waitGroup)
}