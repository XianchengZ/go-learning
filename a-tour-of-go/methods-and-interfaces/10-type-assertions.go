package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// asserts that the inferface value i holds the concrete typt T
	// and assigns the underlying T value to the variable to the left
	s := i.(string)
	fmt.Println(s)
	
	
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
	
	f = i.(float64) // panic
	fmt.Println(f)
}