package main

import (
	"exercise-02-min-max-test/math"
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
