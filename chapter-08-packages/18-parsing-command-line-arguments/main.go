package main

// The `flag` package allows us to parse arguments and flags sent to our program.
import (
	"flag"
	"fmt"
	"math/rand"
)

// Here's an example program that generates a number between 0 and 6.
// We can change the `max` value by sending a flag `-max=100` to the program.
// Run with `go run main.go -max=100`.
func main() {
	// Define flags
	maxp := flag.Int("max", 6, "the max value")
	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}
