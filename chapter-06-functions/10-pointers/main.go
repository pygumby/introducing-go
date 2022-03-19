package main

import "fmt"

// When we call a function that takes an argument, that argument is copied to the function.
// func zero(x int) {
// 	x = 0
// }
// func main() {
// 	x := 5
// 	zero(x)
// 	fmt.Println(x) // `x` is still 5.
// }

// Pointers reference a location in memory where a value is stored rather than the value itself.
// func zero(xPtr *int) { // A pointer is represented by an `*` followed by the type of the value.
// 	*xPtr = 0 // An asterisk is also used to dereference pointer variables.
// }
// func main() {
// 	x := 5
// 	zero(&x)       // We use the `&` operator to find the address of a variable.
// 	fmt.Println(x) // `x` is now 0.
// }

// Another way to get a pointer is to use the built-in new function:
func one(xPtr *int) {
	*xPtr = 1
}
func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // `x` is 1.
}
