package main

import "fmt"

func sliceSum(s []int) (sum int) {
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return
}

func main() {
	fmt.Println(sliceSum([]int{1, 2, 3}))
}
