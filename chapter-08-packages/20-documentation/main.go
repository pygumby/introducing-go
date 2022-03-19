// To view documentation in the terminal, do: `go doc 20-documentation/math Average`
// To view documentation in the browser, do: `godoc -http=localhost:6060`

package main

import (
	"20-documentation/math"
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
