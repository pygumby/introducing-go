package main

import "fmt"

// This won't work:
// func f() {
// 	fmt.Println(x)
// }
// func main() {
// 	x := 5
// 	f()
// }

// But this will work:
// func f(x int) {
// 	fmt.Println(x)
// }
// func main() {
// 	x := 5
// 	f(x)
// }

// And this will work, too:
var x int = 5

func f() {
	fmt.Println(x)
}
func main() {
	f()
}
