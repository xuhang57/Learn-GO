package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Println("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Println("I don't know about the type %T\n", v)
	}
}

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok) // hello, true

	f, ok := i.(float64)
	fmt.Println(f, ok) // hello, false

	f = i.(float64) // panic
	fmt.Println(f)

	do(21)
	do("hello")
	do(true)
}

/*
A type assertion provides access to an interface value's underlying concrete value. t:=i.(T)


*/