package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	fmt.Printf("x is of type %T\n", x)

	// i := 42
	// f := float64(i)
	// u := unit(f)

	// var i int = 42
	// var f float64 = float64(i)
	// var u uint = uint(f)
}

// Unlike in C, in Go assignment between items of different type requires an explicit conversion
