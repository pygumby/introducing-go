package main

import "fmt"

func main() {
	// New map 1: runtime error
	// var x map[string]int
	// x["key"] = 10
	// fmt.Println(x)

	// New map 2: including initialization
	// x := make(map[string]int)
	// x["key"] = 10
	// fmt.Println(x["key"])

	// New map 3: `int` as key
	// x := make(map[int]int)
	// x[1] = 10
	// fmt.Println(x[1])

	// New map 4: Shorthand
	// elements := map[string]string{
	// 	"H": "Hydrogen",
	// 	"He": "Helium",
	// 	"Li": "Lithium",
	// 	"Be": "Beryllium",
	// 	"B":  "Boron",
	// 	"C":  "Carbon",
	// 	"N":  "Nitrogen",
	// 	"O":  "Oxygen",
	// 	"F":  "Fluorine",
	// 	"Ne": "Neon", // Note, again, mandatory comma
	// }

	// delete
	// x := make(map[int]int)
	// x[1] = 10
	// fmt.Println(x[1])
	// delete(x, 1)
	// fmt.Println(x[1])

	// Example program 1
	// elements := make(map[string]string)
	// elements["H"] = "Hydrogen"
	// elements["He"] = "Helium"
	// elements["Li"] = "Lithium"
	// elements["Be"] = "Beryllium"
	// elements["B"] = "Boron"
	// elements["C"] = "Carbon"
	// elements["N"] = "Nitrogen"
	// elements["O"] = "Oxygen"
	// elements["F"] = "Fluorine"
	// elements["Ne"] = "Neon"
	// fmt.Println(elements["Li"])
	// fmt.Println(elements["Un"]) // Looking up a non-existing element returns nothing...
	// name, ok := elements["Un"] // ...hence, Go provides a better way.
	// fmt.Println(name, ok)
	// if name, ok := elements["Un"]; ok { // Idiomatic: If `ok`, run the code.
	// 	fmt.Println(name, ok)
	// }

	// Example program 2
	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
		"C": {
			"name":  "Carbon",
			"state": "solid",
		},
		"N": {
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
