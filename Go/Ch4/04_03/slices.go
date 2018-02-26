package main

import (
	"fmt"
	"sort"
)

func main() {
	// Declaring a slice is just like declarings an array, except you do not put a number in the braces
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	// This is how to remove the first item of the colors slice
	//Fact: the 2nd argument after the colon(len(colors)) is not necessary as that is the default
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	// This is how to remove the last item of the slice
	// Fact, the '0' before the colon is not necessary as that is the default
	colors = append(colors[0 : len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 12
	numbers[1] = 7
	numbers[2] = 1
	numbers[3] = 100
	numbers[4] = 43
	fmt.Println(numbers)

	// This adds an item to the numbers slice, and automatically increases its capacity from 5, to 10
	numbers = append(numbers, 235)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	sort.Ints(numbers)
	fmt.Println(numbers)

}
