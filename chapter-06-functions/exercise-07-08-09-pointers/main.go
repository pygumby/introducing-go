package main

import "fmt"

func main() {
	i := 42

	// 7. How do you get the memory address of a variable?
	ptr := &i
	fmt.Println(ptr) // Something like 0xc0000b0008

	// 8. How do you assign a value to a pointer?
	*ptr = 13
	fmt.Println(*ptr) // 13
	fmt.Println(i)    // 13

	// 9. How do you create a new pointer?
	newPtr := new(int)
	fmt.Println(newPtr)  // Something like 0xc0000b0020
	fmt.Println(*newPtr) // 0
}
