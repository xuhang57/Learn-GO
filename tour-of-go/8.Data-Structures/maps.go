package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// var m = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.68433, -7439967,
// 	},
// 	"Google": Vertex{
// 		37.42202, -122.08408,
// 	},
// }

var m = map[string]Vertex{
	"Bell Labs": {40.68433,
		-74.39967},
	"Google": {37.42202,
		-122.08408},
}

func main() {
	// make function returns a map of the given type, initialized and ready for use
	// m = make(map[string]Vertex)
	// m["Bell Labs"] = Vertex{
	// 	40.68433, -7439967,
	// }
	// fmt.Println(m["Bell Labs"])
	fmt.Println(m)

	test := make(map[string]int)

	test["Answer"] = 42
	fmt.Println("The value:", test["Answer"])
	delete(test, "Answer")

	test["Answer"] = 23
	elem, ok := test["Answer"]

}

/*
A map maps keys to values

The zero value of a map is nil. A nil map has no keys, nor can keys be added.
*/
