package main

func main() {

	cards := newDeck()

	z, y := deal(cards, 3)
	z.print()
	y.print()
}
