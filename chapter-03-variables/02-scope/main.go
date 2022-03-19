package main

import "fmt"

// Variable outside the main function
// The `f` function now has access to the `x` variable.
var x string = "Hello, world."

func main() {
	// Variable inside the main function
	// var x string = "Hello, world."
	fmt.Println(x)
}

func f() {
	fmt.Println(x)
}
