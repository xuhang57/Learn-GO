package main

import "fmt"

// fib is a function that returns a func that returns an int
func fib() func() int {
	f1, f2 := 0, 1
	return func() int {
		f := f1
		f1, f2 = f2, f + f2
		return f
	}
}

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}