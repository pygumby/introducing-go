package main

import "fmt"

func main() {
	// New array 1
	// var x [5]int
	// x[4] = 100
	// fmt.Println(x)

	// New array 2
	// x := [5]float64{98, 93, 77, 82, 83}

	// New array 3
	// x := [5]float64{
	// 	98,
	// 	93,
	// 	77,
	// 	82,
	// 	83, // Note mandatory comma
	// }

	// Example program using arrays 1: without `len`
	// var x [5]float64
	// x[0] = 98
	// x[1] = 93
	// x[2] = 77
	// x[3] = 82
	// x[4] = 83
	// var total float64 = 0
	// for i := 0; i < 5; i++ {
	// 	total += x[i]
	// }
	// fmt.Println(total / 5)

	// Example program using arrays 2: with `len`
	// var x [5]float64
	// x[0] = 98
	// x[1] = 93
	// x[2] = 77
	// x[3] = 82
	// x[4] = 83
	// var total float64 = 0
	// for i := 0; i < len(x); i++ {
	// 	total += x[i]
	// }
	// fmt.Println(total / float64(len(x))) // Note the necessary type conversion

	// Example program using arrays 3: special form of `for`
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total float64 = 0
	for _, value := range x { // `_` replaces `i` to avoid the "unused declarations" error
		total += value
	}
	fmt.Println(total / float64(len(x))) // Note the necessary type conversion
}
