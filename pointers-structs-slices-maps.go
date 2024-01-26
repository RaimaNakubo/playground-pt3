package main

import (
	"fmt"
	"math"
)

// compute calls passed function with parameters 3 & 4
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// function closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum //sum is not accessible inside main()
	}
}

func main() {

	//function values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	fmt.Println()

	//closured function calls
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println( //everything printed from here is a closure
			pos(i),    //values of pos and neg is a closures, bc they are referencing adder() function calls
			neg(-2*i), //adder() is located outside of the for loop scope and main() scope
		)
	}

}
