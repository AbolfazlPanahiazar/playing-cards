package main

func main() {
	cards := newDeck()

	// hand, remainingCards := deal(cards, 15)
	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())

	cards.saveToFile("my_cards")
}