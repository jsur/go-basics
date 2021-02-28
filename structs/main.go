package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	height    float32
	isTrue    bool
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@jim.com",
			zipCode: 94000,
		},
	}
	jim.updateName("jimmy")
	jim.print()
}

/*
	Pointer receiver: This function can be called from either
	the root type (person) or the pointer.
	See example above, jim.updateName modifies original firstName in memory, not like when
	a value is passed to a function where a copy is modified.
	https://tour.golang.org/methods/4
*/
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

/*

Go value types:

int
float
string
bool
struct

Go reference types:

slice
map
channel
pointer
func

*/
