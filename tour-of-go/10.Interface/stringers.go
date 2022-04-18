package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Lucas Xu", 12}
	b := Person{"Luke Hui", 24}
	fmt.Println(a, b)
}

/*

One of the most ubiquitous interfaces is Stringer defined by the fmt package

A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values

*/
