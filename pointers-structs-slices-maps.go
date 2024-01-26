package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

// nil map
var m map[string]Vertex

// map literal
var ml = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// map literal omit type name
var mlo = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {

	//basic map syntax
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{ //key - "Bell Labs", value - Vertex{x,y}
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
	fmt.Println()

	fmt.Println(ml)
	fmt.Println()

	fmt.Println(mlo)
	fmt.Println()

	//mutating maps
	mm := make(map[string]int)

	mm["Answer"] = 42
	fmt.Println("The Value:", mm["Answer"])

	mm["Answer"] = 48
	fmt.Println("The Value:", mm["Answer"])

	delete(mm, "Answer")
	fmt.Println("The Value:", mm["Answer"])

	element, isPresent := mm["Answer"]
	fmt.Println("The Value:", element, "is present?", isPresent)

	mm["Answer"] = 0
	element, isPresent = mm["Answer"]
	fmt.Println("The Value:", element, "is present?", isPresent)
}
