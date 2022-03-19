package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List // The zero value for a `List` is an empty list.
	x.PushBack(1)
	X.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
