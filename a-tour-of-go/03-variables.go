package main

import "fmt"

// variables declared without an explicit initial value are
// given their zero value

// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}