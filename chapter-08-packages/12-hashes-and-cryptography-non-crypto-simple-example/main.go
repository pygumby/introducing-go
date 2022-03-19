package main

// The non-cryptographic hash functions can be found underneath the `hash` package and include
// `adler32`, `crc32`, `crc64` and `fnv`.
// Here's an example using `crc32`:
import (
	"fmt"
	"hash/crc32"
)

func main() {
	// Create a hasher
	h := crc32.NewIEEE()
	// Write our data to it
	h.Write([]byte("test"))
	// Calculate the `crc32` checksum
	v := h.Sum32()
	fmt.Println(v)
}
