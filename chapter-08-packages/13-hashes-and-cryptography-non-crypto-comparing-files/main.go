package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func getHash(fileName string) (uint32, error) {
	// Open the file
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	// Remember to always close opened files
	defer f.Close()
	// Create a hasher
	h := crc32.NewIEEE()
	// Copy the file into the hasher
	_, err = io.Copy(h, f)
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

// A common use for `crc32` is to compare two files. If the `Sum32` value for both files is the
// same, it's highly likely (though not 100% certain) that the files are the same. If the values
// are different, then the files are definitely not the same:
func main() {
	h1, err := getHash("test1.txt")
	if err != nil {
		return
	}
	h2, err := getHash("test2.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}
