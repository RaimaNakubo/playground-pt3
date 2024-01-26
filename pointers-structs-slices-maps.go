/*
Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/

package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	splittedString := strings.Fields(s) //splitted string is a slice of words
	fmt.Println()
	fmt.Println(splittedString)

	ret := make(map[string]int) //declaring a map to return

	for totalWordsCounter := range splittedString { //iterating through every word
		fmt.Println("Counted words:", totalWordsCounter+1)
		ret[splittedString[totalWordsCounter]] += 1 //counting exactly the same words(incrementing the word counter) & filling the map with [word]wordCounter
	}

	fmt.Println("returned map:", ret)
	return ret
}

func main() {
	//wc.Test runs a test and prints success or failure
	wc.Test(WordCount)
}
