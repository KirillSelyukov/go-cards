package main

func main() {
	cards := deck{newCard()}
	cards = append(cards, "Six of Spades")
	cards = append(cards, "Seven of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
