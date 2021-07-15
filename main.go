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

	jimPointer := &jim
	jimPointer.updateFirstName("Abol")
	jim.print()
}

func (p person) print() {
	fmt.Println(p)
}

func (p *person) updateFirstName(fn string) {
	(*p).firstName = fn
}
