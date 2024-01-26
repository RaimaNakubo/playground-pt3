package main

import (
	"fmt"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	//slice append
	var s []int //len 0 cap 0 - nil slice
	printSlice(s)

	s = append(s, 0) //len 1 cap 1 - [0]
	printSlice(s)

	s = append(s, 1) //len 2 cap 2	- [0,1]
	printSlice(s)

	s = append(s, 2, 3, 4) //len 5 cap 6 - [0,1,2,3,4] - cap != 5 bc adding multiple elements
	printSlice(s)

	s = append(s, 5) //len 6 cap 6 - [0,1,2,3,4,5]
	printSlice(s)
	fmt.Println()

	//range
	for i, v := range pow { //range returns two values from pow every iteration
		fmt.Printf("2**%d = %d\n", i, v)
	}
	fmt.Println()

	//skipping value
	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // pow[i] == 2**i
	}
	//skipping index
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// printSlice displaying information about slice
func printSlice(s []int) {
	fmt.Printf("Length: %d Capacity: %d Content: %v\n", len(s), cap(s), s)
}
