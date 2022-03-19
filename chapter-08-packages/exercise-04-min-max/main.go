package main

import (
	"exercise-04-min-max/math"
	"fmt"
)

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(math.Average(xs))
	fmt.Println(math.Min(xs))
	fmt.Println(math.Max(xs))
}
