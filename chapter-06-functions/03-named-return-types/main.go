package main

import "fmt"

func f() (r int) {
	r = 1
	return
}

func main() {
	fmt.Println(f())
}
