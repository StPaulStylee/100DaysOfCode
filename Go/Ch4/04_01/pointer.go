package main

import "fmt"

func main() {
	var p *int

	if p != nil {
		//If you want to get the value that the pointer is pointing to, you must put the '*' in front of it
		fmt.Println("Value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}

	var v int = 42
	// The '&' sets the pointer to the var 'v'
	p = &v

	if p != nil {
		//If you want to get the value that the pointer is pointing to, you must put the '*' in front of it
		fmt.Println("Value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}

	var value1 float64 = 42.13
	// You can set pointers implicity
	pointer1 := &value1
	fmt.Println("Value 1:", *pointer1)

	// You can edit the value of the pointer and hence, the value of the value stored in memory
	*pointer1 = *pointer1 / 31
	fmt.Println("Value 1:", *pointer1)
	fmt.Println("Value 1:", value1)

}
