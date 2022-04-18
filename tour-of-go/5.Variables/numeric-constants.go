package main

import "fmt"

const (
	// Big creates a huge number by shifting a 1 bit left 100 places
	Big = 1 << 100
	// Small shifts it right again 99 places, so we end up with 1 << 1, or 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func constant() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// fmt.Println(needInt(Big))
	// ./prog.go:19:21: constant 1267650600228229401496703205376 overflows int

}
