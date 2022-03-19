package main

import "fmt"

func main() {
	// New constant: One statement
	const x string = "Hello, world."
	fmt.Println(x)

	// Of course, type inference works here as well.

	// Compile-time error: "cannot assign to x (declared const)"
	// x = "some other string"
}
