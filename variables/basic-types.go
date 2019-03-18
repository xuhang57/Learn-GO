package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// bool, string, int/uint, int8/uint8, int16/uint16, int32/uint32, int64/uint64, uintptr
// byte (alias for uint8), rune (alias for int32, represents a Unicode code point)
// float32, float64, complex64, complex128
// int/uint uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
// Variables declared without an explicit initial value are given their zero value.
