package main

import (
	"testing"
	"sync"
)

func Test_updateMessage(t *testing.T){
	msg = "hello world"

	var mutex sync.Mutex

	waitGroup.Add(2)
	go updateMessage("Goodbye", &mutex)
	go updateMessage("X", &mutex)
	waitGroup.Wait()

	if msg != "Goodbye" {
		t.Error("incorrect value in msg")
	}
}