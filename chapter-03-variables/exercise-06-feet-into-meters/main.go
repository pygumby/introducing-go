package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Enter a distance in ft: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 0.3048

	fmt.Println(math.Round(output*100)/100, "m")
}
