package main

import "fmt"

func fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	fmt.Println(fib(42)) // 267914296
}
