package main

import "fmt"

// func add( x int, y int ) int {
// 	return x + y
// }

// alternative same function declaration
func add(x, y int) int {
	return x + y
}

// multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// named return value
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(20, 30))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}