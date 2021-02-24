package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	height    float32
	isTrue    bool
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anonymous"}
	var alex person // zero value struct assigned to alex
	fmt.Println(alex)
	// %+v prints out all field names and values from alex
	fmt.Printf("%+v", alex)
	alex.firstName = "John"
	fmt.Printf("%+v", alex)
}
