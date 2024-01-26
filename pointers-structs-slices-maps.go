package main

import (
	"fmt"
	"strings"
)

// basic struct syntax
type Vertex struct {
	X int
	Y int
}

// struct literals
var (
	v1  = Vertex{1, 2}  //v1 has type Vertex
	v2  = Vertex{X: 1}  //v2.X == 1, v2.Y == 0 by default(implicit)
	v3  = Vertex{}      //v3.X == 0, v3.Y == 0
	pV1 = &Vertex{1, 2} //pV1 has type *Vertex
)

func main() {

	//pointers basic syntax
	i, j := 42, 2701

	p := &i         //pointer to i
	fmt.Println(*p) //reading i through the p pointer

	*p = 21        //writing i through the p pointer
	fmt.Println(i) //reading i's new value

	p = &j         //pointer to j
	*p = *p / 37   //dividing j through the p pointer
	fmt.Println(j) //reading j's new value
	fmt.Println()

	//creating a struct
	fmt.Println(Vertex{1, 2})

	//accessing X field of a v struct
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	//accessing X field of a v struct by reference
	pV := &v
	pV.X = 1e9
	fmt.Println(v)
	fmt.Println()

	//structs initialized by literals
	fmt.Println(v1, pV1, v2, v3)
	fmt.Println()

	//basic array syntax
	var a [2]string
	a[0] = "Job"
	a[1] = "offer"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	array := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(array)
	fmt.Println()

	//basic slice syntax
	var slice []int = array[1:4]
	fmt.Println(slice)
	fmt.Println()

	//referencing slices
	names := [4]string{
		"Pyotr",
		"Nikolai",
		"Alexey",
		"Kirill",
	}
	fmt.Println(names)

	//b and c represents a subset of names
	b := names[0:2]
	c := names[1:3]
	fmt.Println(b, c)

	c[0] = "HTO?" //changing a subset
	fmt.Println(b, c)
	fmt.Println(names)
	fmt.Println()

	//slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	//slice defaults
	q = q[1:4]
	fmt.Println(q)

	q = q[:2]
	fmt.Println(q)

	q = q[1:]
	fmt.Println(q)
	fmt.Println()

	//length and capacity of slices
	w := []int{2, 3, 5, 7, 11, 13}
	printSlice("", w)

	fmt.Println("re-slicing resulting zero length:")
	w = w[:0]
	printSlice("", w)

	fmt.Println("Extending length by 4:")
	w = w[:4]
	printSlice("", w)

	fmt.Println("Removing first two elements:")
	w = w[2:]
	printSlice("", w)

	/*
		fmt.Println("Extending length beyond slice's capacity:")
		w = w[:5] //"runtime error: slice bounds out of range [:5] with capacity 4"
		printSlice(w)
	*/

	//nil slices
	var ss []int
	fmt.Println(ss, len(ss), cap(ss))
	if ss == nil {
		fmt.Println("This is a nil slice!")
	}
	fmt.Println()

	//creating a slice with make
	ma := make([]int, 5) //len 5 cap 5
	printSlice("ma", ma)

	mb := make([]int, 0, 5) //len 0 cap 5
	printSlice("mb", mb)

	mc := mb[:2] //len 2 cap 5
	printSlice("mc", mc)

	md := mc[2:5] //len 3 cap 3
	printSlice("md", md)

	//slices of slices
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func printSlice(s string, x []int) {
	fmt.Println("Slice's dimensions:")
	fmt.Printf("%s length: %d capacity: %d content: %v\n", s, len(x), cap(x), x)
	fmt.Println()
}
