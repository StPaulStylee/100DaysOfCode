package main

import (
	"fmt"
)

// The defer keyword executes everything in the function first, then runs the defer statement
// If you have multiple defer statements, they run in LIFO order
func main() {
	defer fmt.Println("Close the file")
	fmt.Println("Open the file")

	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")

	myFunc()

	defer fmt.Println("Statement 3")
	defer fmt.Println("Statement 4")
	fmt.Print("Undeferred statement")

}

func myFunc() {
	defer fmt.Println("Deferred myFunction")
	fmt.Println("Not deferred in myFunction")
}
