package main

import "os"

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		// Handle error here
		return
	}
	defer file.Close()
	file.WriteString("Hello, world.\n")
}
