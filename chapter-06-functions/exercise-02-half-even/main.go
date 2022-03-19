package main

import "fmt"

func halfEven(i int) (half int, even bool) {
	half = i / 2
	even = i%2 == 0
	return
}

func main() {
	half, even := halfEven(2)
	fmt.Println(half, even)
}
