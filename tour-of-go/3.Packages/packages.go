package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// The environment in which these programs are executed is deterministic, so each
	// time you run the example program rand.Intn will return the same number
	// To see a different number, seed the number generator; see rand.Seed
	fmt.Println("My favorite number is", rand.Intn(10))
}
