package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Abolfazl",
		lastName:  "Panahiazar",
		contactInfo: contactInfo{
			email:   "abolfazlpanahiazar@gmail.com",
			zipCode: 123456,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateFirstName("Abol")
	// or
	(&jim).updateFirstName("Abolf")
	jim.print()
}

func (p person) print() {
	fmt.Println(p)
}

func (p *person) updateFirstName(fn string) {
	// (*p).firstName = fn
	// or
	p.firstName = fn
}

// Value types
// int, float, string, bool, struct

// Refrence types (don't worry about pointers with these)
// slides, maps, channels, pointers, functions
