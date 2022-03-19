package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// We could also sort by age by doing this:
type ByAge []Person

func (ps ByAge) Len() int {
	return len(ps)
}
func (ps ByAge) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}
func (ps ByAge) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}

	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}
