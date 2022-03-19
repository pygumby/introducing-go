package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// This is precisely how the fmt.Println function is implemented:
// `func Println(a ...interface{}) (n int, err error)`

func main() {
	// Call to variadic function 1
	// fmt.Println(add(1, 2, 3))

	// Call to variadic function 2
	xs := []int{1, 2, 3}
	fmt.Println(add(xs...)) // Note the three dots
}
