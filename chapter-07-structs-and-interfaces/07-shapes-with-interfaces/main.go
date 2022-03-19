package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// Go has a way of making accidental similarities explicit through a type known as an interface.
// Like a struct, an interface is created using the `type` keyword.
type Shape interface {
	// Instead of defining fields, we define a method set.
	// A method set is a list of methods that a type must have to implement the interface.
	// Both `Rectangle` and `Circle` have `area` methods that return `float64`s.
	// Therefore, both types implement the `Shape` interface.
	// Nothing additional is required to implement an interface.
	// There is no `implements` or `extends` keyword.
	area() float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

// See https://www.oreilly.com/catalog/errata.csp?isbn=0636920046516, page 61, chapter 7
func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

type Circle struct {
	x, y, r float64
}

// See https://www.oreilly.com/catalog/errata.csp?isbn=0636920046516, page 61, chapter 7
func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Suppose we want to write a function that calculates the area of several shapes.
// This is the problem interfaces are designed to solve.
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// Interfaces can also be used as fields.
// Consider a `MultiShape` that is made up of several smaller shapes:
type MultiShape struct {
	shapes []Shape
}

// We can even turn `MultiShape` itself into a `Shape` by giving it an `area` method:
func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.area())

	// We would call this function like this:
	fmt.Println(totalArea(&c, &r))

	// We can create a MultiShape like this:
	multiShape := MultiShape{
		shapes: []Shape{
			Circle{0, 0, 5},
			Rectangle{0, 0, 10, 10},
		},
	}
	fmt.Println(multiShape.area())
}
