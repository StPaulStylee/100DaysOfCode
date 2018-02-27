package main

import "fmt"

func main() {
	sum := 1
	fmt.Println("Sum", sum)

	colors := []string{"Red", "Green", "Blue"}

	// There is the traditional C style for loop available in Go but there is also this other one....
	// This is saying to set the value of i to the current index of the colors slice
	for i := range colors {
		fmt.Println(colors[i])
	}

	// There is no 'while' loop, but you can still do it this way...
	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
	}

	// Further you can break out of loops in various ways
	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 200 {
			goto endofprogram
		}
		if sum > 500 {
			break
		}
	}

endofprogram:
	fmt.Println("end of program")
}
