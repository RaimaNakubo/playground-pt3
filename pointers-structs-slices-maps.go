package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacchi() func() int {
	n1 := 0
	n2 := 1
	return func() int {
		n1, n2 = n2, n1+n2
		return n2 - n1
	}
}

func main() {
	f := fibonacchi() //f is the instance of fibonacchi(); it returns a function with pre-set parameters, not a function call
	for i := 0; i < 10; i++ {
		fmt.Println(f()) //every time f() is called, parameters are NOT re-set to the beginning values bc fibonacchi call is not happening
		//calling f() changes the parameters and returns resulting value
	}
	fmt.Println()

	//here is what it looks like if fibonacchi() is called instead
	/*
		for i := 0; i < 10; i++ {
			fmt.Printf(fibonacchi()) // fmt.Println arg fibonacchi() is a func value, not called
		}
	*/

	//if f() is called over and over, parameters are still not reset
	for j := 0; j < 10; j++ {
		fmt.Println(f())
	}
	fmt.Println()

	//however if fibonacchi() is called again it will re-set the parameters; calculations will start from the beginning
	f = fibonacchi()
	for k := 0; k < 10; k++ {
		fmt.Println(f())
	}
	fmt.Println()

}
