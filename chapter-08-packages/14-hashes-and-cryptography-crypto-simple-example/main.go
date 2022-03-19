package main

import (
	"crypto/sha1"
	"fmt"
)

// Cryptographic hash functions are similar to their non-cryptographic counterparts, but they have
// the added property of being hard to reverse. One common cryptographic hash function is known as
// SHA-1.
// Here's how it is used:
func main() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
