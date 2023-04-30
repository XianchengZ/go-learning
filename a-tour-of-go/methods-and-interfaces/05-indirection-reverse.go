package main

import (
	"fmt"
	"math"
)


type Vertex struct {
	X, Y float64
}

// methods with pointer receiver can take either a value or pointer
// methods with value receiver can take either a value or pointer
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// functions with pointer receiver can only take pointer
// functions with value receiver can only take value
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}



func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}