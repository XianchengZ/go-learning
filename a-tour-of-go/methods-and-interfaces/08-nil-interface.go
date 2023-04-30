package main

import "fmt"

type I interface {
	M()
}

// runtime error
// no type inside the interface tuple to indicate which concrete method to call
func main() {
	var i I
	describe(i)
	i.M()
}


func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

