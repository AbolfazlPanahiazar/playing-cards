package main

import "fmt"

type deck []string

// print receive function of deck type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}