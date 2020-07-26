package main

import (
	"fmt"
	"math"
)

// Vertex is a type contains two floats
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

function ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
	ScaleFunc(&v, 10)
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println(v. p)
	fmt.Println(AbsFunc(*p))


}

/*
You can declare methods with pointer receivers. Type can be *T for some type T.

Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often
need to modify their receiver, pointer receivers are more common than value receivers.

Functions with a pointer argument must take a pointer. However, methods with pointer receivers take either a value of a pointer
as the receiver when they are called.

Functions that take a value argument must take a value of that specific type. However, methods with value receivers take either a
value of a pointer as the receiver when they are called

Two reasons to use a pointer receiver in a method:
1. The method can modify the value that its receiver points to.
2. Avoid copying the value on each method call. More efficient if the receiver is a large struct

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

*/