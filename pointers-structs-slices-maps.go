//slice exercise

/*
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.\
When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)
*/

package main

import (
	"golang.org/x/tour/pic"
)

// Pic creates the shape of an image based on a function (a string to be displayed at go playground website)
func Pic(dx, dy int) [][]uint8 {
	pictureSlice := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		partialPictureSlice := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			partialPictureSlice[x] = uint8((x^y)*((x+y)/2) + (x * y)) //change this function
		}
		pictureSlice[y] = partialPictureSlice
	}

	return pictureSlice
}

func main() {
	pic.Show(Pic)
}
