package main

import (
	"fmt"
	"strings"
)

type deck []string

// print receive function of deck type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i + 1, card)
	}
}

// new deck constructor function
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

// deal function the slices a deck into two deck of cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// deck into string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}