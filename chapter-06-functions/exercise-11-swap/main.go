package main

import "fmt"

func swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

func main() {
	i := 1
	j := 2
	fmt.Println("i:", i, "| j:", j)
	swap(&i, &j)
	fmt.Println("i:", i, "| j:", j)
}
