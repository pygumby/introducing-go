package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Enter a temperature in °F: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5 / 9

	fmt.Println(math.Round(output*10)/10, "°C")
}
