package main

import "fmt"

func main() {
	// Shorthand to define multiple variables
	var ( // or `const`
		a = 5
		// b = 10 // Commented out to avoid annoying "declared but not used" error
		// c = 15
	)
	fmt.Println("Value of a:", a)
}
