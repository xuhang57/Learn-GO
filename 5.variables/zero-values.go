package main

import "fmt"

func zeros() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println("%v %v %v %q\n", i, f, b, s)
}

// Variable declared without an explicit initial value are given their zero value.
/*
Zero value is:
  0 for numeric types,
  false for the boolean type, and
  "" (the empty string) for strings

*/
