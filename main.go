package main

func main() {

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.print()
	cards.shuffle()
	cards.print()
	// hand, rest := deal(cards, 3)
	// hand.print()
	// rest.print()

}
