package main

import (
	"testing"
	"os"
	"sync"
	"io"
	"strings"
)

func Test_printSomething(t *testing.T){
	stdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go printSomething("epsilon", &waitGroup)

	waitGroup.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdout

	if ! strings.Contains( output, "epsilon"){
		t.Errorf("Expected to find epsilon, but it is not there")
	}
}