package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	fmt.Println(a.Abs())
	// v is a Vertex not *Vertex, does not implement Abser
	a = v
	fmt.Println(a.Abs())

	var i I = T{"Hello"}
	i.M()

	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var j I
	describe(j)
	j.M()

	var w interface{}
	describeInterface(w)

	w = 42
	describeInterface(w)

	w = "hello"
	describeInterface(w)
}

type Myfloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

// This method meains type T implements the interface I
// But we don't need to explicitly declare that it does so
func (t T) M() {
	fmt.Println(t.S)
}

func (f F) M() {
	fmt.Println(f)
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeInterface(i interface{}) {
	fmt.Printf("(%v, %T", i, i)
}

/*
An interface type is a defined as a set of method signatures. A value of interface type can hold any value that
implements those method

A type implements an interface by implementing its methods. There is no explicit declaration of intent. Implicit
interfaces decouple the definition of an interface from its implemention, which could then appear in any package.


Under the hood, interface values can be thought of as tuple of a value and a  concrete type: (value, type)

An interface value holds a value of a specific underlying concrete type. Calling a method on an interface value
executes the method of the same name on its underlying type.

If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

Note that an interface value that holds a nil concrete value is itself non-nil.

A nil interface value holds neither value nor concrete type. Calling a method on a nil interface is a run-time error
because there is no type inside the interface to indicate which conrete method to call.

The interface type that specifies zero methods is known as the empty interface: interface{}. An empty interface may
hold values of any type. (Every type implements at least zero methods)
*/
