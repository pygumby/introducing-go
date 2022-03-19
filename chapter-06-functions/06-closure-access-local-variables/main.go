package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))

	// When you create a local function like this, it also has access to other local variables.
	// In this case, `increment` and the variable `x` form the closure.
	x := 10
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
