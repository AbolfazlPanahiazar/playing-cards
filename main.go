package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person
	fmt.Println(alex)
}
