package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/* Abs method has a receiver of type Vertex named v */
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

/*
Go does not have classes. However, you can define methods on types.package  Methods

A method is a function with a special receiver argument.

*/