package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// methods with pointer receiver can take either a value or pointer
// methods with value receiver can take either a value or pointer
func (v *Vertex) Scale(value float64){
	v.X = v.X * value
	v.Y = v.Y * value
}

// functions with pointer receiver can only take pointer
// functions with value receiver can only take value
func ScaleFunction(v * Vertex, value float64){
	v.X = v.X * value
	v.Y = v.Y * value
}

func main(){
	var v = Vertex{3, 4}

	fmt.Printf("Origin value v: %v\n", v)
	v.Scale(2)
	fmt.Println(v)
	
	// ScaleFunction(v, 10) // Compile Error
	ScaleFunction(&v, 10) // OK
	fmt.Println(v)

	p := &Vertex{4, 3} // p is a pointer

	p.Scale(3) // implicitly (*p).Scale(3)
	fmt.Println(p)

	ScaleFunction(p, 10)
	fmt.Println(p)
}