package main

import "fmt"

func main() {
	// New slice 1
	// var x []float64
	// x = make([]float64, 5) // 5 represents the length of the underlying array.

	// New slice 2
	// var x []float64
	// x = make([]float64, 5, 10) // 10 represents the capacity of the underlying array.

	// New slice 3
	// arr := [5]float64{1, 2, 3, 4, 5}
	// x := arr[0:5] // also valid: `arr[0:]`, `arr[:5]` and `arr[:]`

	// append
	// slice1 := []int{1, 2, 3}
	// slice2 := append(slice1, 4, 5)
	// fmt.Println(slice1, slice2)

	// copy
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
