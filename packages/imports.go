package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))

	// In Go, a name is exported if it begins with a capital letter.
	fmt.Println(math.Pi)
}
