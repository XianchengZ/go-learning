package main

import "fmt"

func main() {
	sum := 0
	// for loop
	for i := 0; i < 10; i++ {
		sum += i
	}

	// while loop
	for sum < 1000 {
		sum += sum
	}

	// forever loop
	// for{

	// }

	fmt.Println(sum)
}