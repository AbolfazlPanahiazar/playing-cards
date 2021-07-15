package main

import "fmt"

type contactInfo struct {
	email string
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
		lastName: "Panahiazar",
		contactInfo: contactInfo{
			email: "abolfazlpanahiazar@gmail.com",
			zipCode: 123456,
		},
	}
	fmt.Println(jim)
}
