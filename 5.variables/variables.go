package main

import "fmt"

// var c, python, java bool
var i, j int = 1, 2

func main() {
	// var i int
	// fmt.Println(i, c, python, java)

	// var c, python, java = true, false, "no!"
	// fmt.Println(i, j, c, python, java)

	// Inside a function, := shorts for assignment statement can be used in place of a var
	// declaration with implicit type.
	// Outside a function, every statement begins with a keyword(var, func, etc.) and so
	// the := construct is not available
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
