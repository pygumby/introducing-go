package main

import "fmt"

func greatest(nums ...int) int {
	greatest := nums[1]
	for i := 1; i < len(nums); i++ {
		if nums[i] > greatest {
			greatest = nums[i]
		}
	}
	return greatest
}

func main() {
	fmt.Println(greatest(1, 2, 3, 4, 3, 2, 1))
}
