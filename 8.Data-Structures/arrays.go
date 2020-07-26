package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// a slice does not store any data, it just describes a section of an underlying array
	// changing the elements of a slice modifies the corresponding elements of its underlying array
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	c := names[0:2]
	d := names[1:3]
	fmt.Println(c, d)

	d[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// A slice literal is like an array literal without the length
	// builds a slice that references it
	t := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(t)

	l := []int{2, 3, 5, 7, 11, 13}
	l = l[1:4]
	fmt.Println(l)
	l = l[:2]
	fmt.Println(l)
	l = l[1:]
	fmt.Println(l)

	z := []int{2, 3, 5, 7, 11, 13}
	z = z[:0]
	printSlice(z) // len=0, cap=6

	z = z[:4]
	printSlice(z) // len=4, cap=6

	z = z[2:]
	printSlice(z) // len=2, cap=4

	/* The zero value of a slice is nil, which has a len and cap of 0, and no underlying array*/
	var k []int
	fmt.Println(s, len(k), cap(k))
	if k == nil {
		fmt.Println("nil!")
	}

	/*  Slices can be created with the built-in make function */
	/* This is how you create dynamically-size arrays */

	sample := make([]int, 5)   // len(a) = 5
	slice := make([]int, 0, 5) // len(b) = 0, cap(b) = 5
	s = s[1:]                  // len(b) = 4, cap(b) = 4
	
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	b = b[:cap(b)] // len(b) = 5, cap(b)=5
	b = b[1:] // len(b)=4, cap(b)=4

	/* Slices of slices */

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"

	var app []int
	// func append(s []T, vs ... T) []T
	app = append(app, 0)
	app = append(app, 2, 3, 4)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

var pow = []int{1,2,4,8,16,32,64,128}
func rangeExample() {
	for i, v := range pow {
		fmt.Println("2**%d = %d\n", i, v);
	}
}

/* You can skip the index or value by assigning to "_" */
/* Or just omit the second variable to use only index */
func rangeCont(pow []int) {
	for i := range pow {
		pow[i] = 1 << uint(i) // 2**i
	}
}

// array's length is part of its type, so arrays cannot be resized
// Go provides a convenient way of working with arrays
// a[low: high], incluses the first element, but excludes the last one.
// A slice has both a length and a capacity. The length of a slice is the number of elements it contains
// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice
// Appending: If the backing array of s (first argument) is too small to fit all the given values a bigger array will be allocated
// The returned slice will point to the newly allocated array