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

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

type Circle struct {
	x, y, r float64
}

// Between the keyword `func` and the name of the function, we've added a receiver.
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	c := Circle{x: 0, y: 0, r: 5}
	// The receiver allows us to call the function using the `.` operator.
	// We no longer need the `&` operator, as Go automatically knows to pass a pointer.
	fmt.Println(c.area())
}
