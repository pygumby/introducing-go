package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		// Handle error here
		return
	}

	// We use defer `file.Close()` right after opening the file to make sure the file is closed
	// as soon as the function completes.
	defer file.Close()

	// Get file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// Read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
