package main

import "fmt"

// defer
// func first() {
// 	fmt.Println("1st")
// }
// func second() {
// 	fmt.Println("2nd")
// }
// func main() {
// 	defer second()
// 	first()
// }
// `defer` is often used when resources need to be freed in some way.
// f, _ := os.Open(filename)
// defer f.Close()

// panic
// func main() {
// 	panic("PANIC!")
// 	str := recover() // This will never happen!
// 	fmt.Println(str) // `recover`ing from a `panic` doesn't work like this.
// }

// recover
func main() {
	defer func() {
		str := recover()
		fmt.Println(str, "-- but recovered from it!")
	}()
	panic("PANIC")
}
