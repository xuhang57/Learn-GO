package main

import (
	"fmt"
	"math"
)

// Vertex contains two floats: X and Y
type Vertex struct {
	X, Y float64
}

// MyFloat is a simple float type
type MyFloat float64

// Abs method has a receiver of type Vertex named v
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Abs function with no change in functionality
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Abs is a method on a non-struct type
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	// func
	fmt.Println(Abs(v))
}

/*
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument. The receiver appears in its own argument list between the
func keyword and the method name.

You can only declare a method with a receiver whose type is defined in the same package as the method.

*/
