package main

import (
	"fmt"
	"strings"
)

func main() {
	strl1 := "An implicity typed string."
	fmt.Println(strl1)
	// These %x thinga are called verbs. %v prints the value and %T prints the data type.
	// The argument must be passed twice, once for each verb
	fmt.Printf("strl1: %v:%T\n", strl1, strl1)

	var strl2 string = "An explicity typed string."
	fmt.Printf("strl2: %v:%T\n", strl2, strl2)

	fmt.Println(strings.ToUpper(strl1))
	fmt.Println(strings.Title(strl1))

	lValue := "hello"
	uValue := "HELLO"
	fmt.Println("Equal?", (lValue == uValue))
	// strings.EqualFold() is the same as String.equalsIgnoreCase() in Java
	fmt.Println("Equals non-case sensitive?", strings.EqualFold(lValue, uValue))

	fmt.Println("Contains 'exp'?", strings.Contains(strl1, "exp"))
	fmt.Println("Contains 'exp'?", strings.Contains(strl2, "exp"))

}
