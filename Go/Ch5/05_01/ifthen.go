package main

import (
	"fmt"
)

func main() {

	var x float64 = 42
	var result string

	if x < 0 {
		result = "Less than zero"
	} else {
		result = "Greater than or equal to zero"
	}

	fmt.Println("Result:", result)

	//This is a variation of the if statement that only makes the value of x available DURING the if statement
	// execution. Once the 'if' is over, garbage collection collects deletes the value and x is undefined
	// if x := 42; x < 0 {

	// }

}
