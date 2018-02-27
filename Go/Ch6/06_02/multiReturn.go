package main

import (
	"fmt"
)

func main() {
	n1, l1 := FullName("Jeffrey", "Miller")
	fmt.Printf("Fullname: %v, number of char: %v\n", n1, l1)

	n2, l2 := FullNameNakedReturn("Geoffery", "Mueller")
	fmt.Printf("Fullname: %v, number of chars: %v\n", n2, l2)
}

// This function is "getting" the full name, but in Go is is a convention to not use the term get,
// but to simply just name the function the thing that it is returning
func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

// This is another way to write the abovie function, you declare the variable names in the function definition
// and then when you return the values, you don't have to reference the variable names
// also notice that you aren't using implicit typing when the variable values are set, you just use '=' because their types are set in the def
func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
