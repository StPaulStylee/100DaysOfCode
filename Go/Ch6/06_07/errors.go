package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.ext")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err)
	}

	myError := errors.New("My error string")
	fmt.Println(myError)

	attendance := map[string]bool{
		"Ann":  true,
		"Mike": true} //Apparently, at the end of a map declaration you have to put the '}' right there

	// the 'ok' can be anything, but it is Go convention for it to be stated as true
	attended, ok := attendance["M"]
	if ok {
		fmt.Println("Mike attended:", attended)
	} else {
		fmt.Println("No info for Mike")
	}
}
