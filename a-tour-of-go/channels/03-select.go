package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// select statement lets a goroutine wait on
		// multiple communication operations
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// c and quit are two channels shared by two threads/goRoutines
	c := make(chan int)
	quit := make(chan int)
	
	// another thread / goRoutine
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
			time.Sleep(2 * time.Second)
		}
		quit <- 0
	}()

	// fibonacci inside main's thread
	fibonacci(c, quit)
}
