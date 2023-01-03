// packge main is a special package
// it allows go to create an executable file
package main

// same as printf in C
import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	// exported -> first letter is uppercase
	fmt.Println("hello world!")	
	fmt.Println(quote.Go())	
}