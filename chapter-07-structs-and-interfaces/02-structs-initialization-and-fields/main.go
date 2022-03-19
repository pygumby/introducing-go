package main

import "fmt"

// Structs
// A `struct` is a type that contains named fields.
// type Circle struct {
// 	x float64
// 	y float64
// 	r float64
// }
// Shorthand:
type Circle struct {
	x, y, r float64
}

func main() {
	// Initialization
	// This will create a local `Circle` variable that is by default set to zero.
	// For a struct, zero means each of the fields is set to their corresponding zero value.
	// var c1 Circle
	// We can also use the `new function, which returns a pointer to the struct. (`*Circle`)
	// Using `new` with "zero initialization" is somewhat uncommon, however.
	// c2 := new(Circle)
	// Instead, we want to give each of the fields a special value.
	// Option 1, creates a `Circle` variable:
	// c3 := Circle{x: 0, y: 0, r: 5}
	// Option 2, also creates a `Circle` variable:
	// c4 := Circle{0, 0, 5}
	// Option 2, creates a pointer to the struct:
	// c5 := &Circle{0, 0, 5}

	// Fields
	// We can access fields using the `.` operator.
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.x, c.y, c.r)
	c.x = 10
	c.y = 5
	fmt.Println(c)
}
