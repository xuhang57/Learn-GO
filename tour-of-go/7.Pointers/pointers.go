package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer

	*p = 21        // set i through the pointer
	fmt.Println(i) // see new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j
	fmt.Println(j) // see the new value of j
}

// type *T is a pointer to a T value. Its zero value is nil
// var p *int
// Unlike C, Go has no arithmetic

// & operator generates a pointer to its operand
// * operator denotes the pointer's underlying value