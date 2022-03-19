package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ") // Note, not `Println`
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
