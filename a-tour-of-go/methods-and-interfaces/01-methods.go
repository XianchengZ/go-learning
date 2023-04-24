// A method is a function with a special receiver argument.

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// v is the receiver for the Abs function below
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y )
}

// you can declare a method on non-struct types too
type MyFloat float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// You can only delcare a method with a receiver whose type
// is defined in the same package as the method.
// You cannot declare a method with a receiver whose type is defined
// in another package

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}