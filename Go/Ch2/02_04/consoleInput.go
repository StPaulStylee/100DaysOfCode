package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var s string
	// In the Scanln() you pass in a 'reference' of the string, hence the '&s'
	// The Scan function is designed to also parse a string, so it will break up a string input
	// wherever there is a space. If you input "one two theree" the output will only be "one"
	// fmt.Scanln(&s)
	// fmt.Println(s)

	// To simply collect user input from the console... the "os.Stdin" arg refers to input collection via the console
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")         // Notice, just the Print()
	str, _ := reader.ReadString('\n') // the single quote infers a byte value
	fmt.Println(str)

	// If you want to have a number input you ALMOST do the exact some thing
	fmt.Print("Enter a number: ")                            // Notice, just the Print()
	str, _ = reader.ReadString('\n')                         // the single quote infers a byte value
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64) // Parse a string into a float, trimming any whitespace from string prior to conversion
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", f)
	}
}
