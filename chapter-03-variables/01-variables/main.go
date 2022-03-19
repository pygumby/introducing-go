package main

import "fmt"

func main() {
	// New variable 1: One statement
	// var x string = "Hello, world."
	// fmt.Println(x)

	// New variable 2: Two statements
	// var x string
	// x = "Hello, world."
	// fmt.Println(x)

	// Intermezzo: What are declared variables initialized with?
	// If a variable is not assigned any value, Go automatically initializes it with the zero
	// value of the variable's type. E.g., `0` is the zero value of `int`.
	// - https://golangbot.com/variables/
	// var x string
	// fmt.Println("Value of x: ", x)

	// New variable 3: Type inference
	// var x = "Hello, world."
	// fmt.Println(x)

	// New variable 4: Idiomatic shorthand
	// x := "Hello, world."
	// fmt.Println(x)

	// Reassignment
	// var x string
	// x = "first "
	// fmt.Println(x)
	// x = x + "second" // Shorthand: `x += "second"`
	// fmt.Println(x)

	// Equality 1: `false`
	// var x string = "hello"
	// var y string = "world"
	// fmt.Println(x == y)

	// Equality 1: `true`
	// var x string = "hello"
	// var y string = "hello"
	// fmt.Println(x == y)
}
