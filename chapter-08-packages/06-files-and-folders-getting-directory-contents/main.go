package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")
	if err != nil {
		// Handle error here
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1) // By passing in `-1, we return all of the entries.
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
