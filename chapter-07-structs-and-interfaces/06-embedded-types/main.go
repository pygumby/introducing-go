package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// Now, suppose we wanted to create a new `Android` struct.
// This would work, but we would rather say an android *is* a person, rather than *has* one:
// type Android struct {
// 	Person Person
// 	Model  string
// }

// Go supports "is-a" relationships by using embedded types, also referred to as anonymous fields.
// We use the type `Person` and don't give it a name.
type Android struct {
	Person
	Model string
}

func main() {
	p := Person{Name: "Lucas"}
	p.Talk()

	// The `Person` struct can be accessed using the type name:
	a := new(Android)
	a.Person.Talk()
	// ...but we can also call any `Person` methods directly on the `Android`:
	a.Talk()
}
