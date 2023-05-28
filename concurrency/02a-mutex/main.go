package main

import (
	"fmt"
	"sync"
)

var msg string

var waitGroup sync.WaitGroup

// func updateMessage(s string){
// 	defer waitGroup.Done()
// 	msg = s
// }

// func main() {
// 	msg = "Hello, World!"

// 	waitGroup.Add(2)

// 	go updateMessage("Hello, update 1")
// 	go updateMessage("Hello, update 2")

// 	waitGroup.Wait()

// 	fmt.Println(msg)

// }

func updateMessage(s string, mutex *sync.Mutex){
	defer waitGroup.Done()

	mutex.Lock()
	msg = s
	mutex.Unlock()
}

func main() {
	msg = "Hello, World!"

	var mutex sync.Mutex

	waitGroup.Add(2)

	go updateMessage("Hello, update 1", &mutex)
	go updateMessage("Hello, update 2", &mutex)

	waitGroup.Wait()

	fmt.Println(msg)

}