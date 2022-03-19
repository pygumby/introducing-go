package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		// Handle error here
		return
	}
	str := string(bs)
	fmt.Println(str)
}
