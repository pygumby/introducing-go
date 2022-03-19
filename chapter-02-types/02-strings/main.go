package main

import "fmt"

func main() {
	fmt.Println(len("Hello, world."))
	fmt.Println("Hello, world."[1]) // returned as a byte
	fmt.Println("Hello, " + "world.")
}
