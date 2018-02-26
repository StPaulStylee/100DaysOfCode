package main

import "fmt"

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	fmt.Println(str1, str2, str3)

	// Functions can return multiple things in Go. Here, the Println() returns an int and an error object
	// the ':=' is inferred typing of the variables
	stringLength, err := fmt.Println(str1, str2, str3)

	// In Go, if you declare a variable but never use it, an error will be thrown during compilation. Say you didn't want to use the err object
	// Do this... use an '_' and that way you don't have to address the err object
	// stringLength, _ := fmt.Println(str1, str2, str3)

	if err == nil {
		fmt.Println("String length:", stringLength)
	}

	//Below I am using the %v to dynamically print variables
	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of aNumber: %v\n", isTrue)
	// Here, we use %.2f to create a 2 decimal float of our aNumber int, notice the int must be converted to a float!
	fmt.Printf("Value of aNumber: %.2f\n", float64(aNumber))
	// This allows me to print the data types of all the listed variables
	fmt.Printf("Data types: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue)
	// Here I am converting all of those data types into a string and saving the value, fmt.Sprintf returns the string
	myString := fmt.Sprintf("Data types as var: %T, %T, %T, %T, and %T", str1, str2, str3, aNumber, isTrue)
	// You've always gotta use your variables, so I print it
	fmt.Println(myString)

}
