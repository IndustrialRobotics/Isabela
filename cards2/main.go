package main

func main() {

	cards := newCard()
	//cards.print()

	hand, remainingDeck := deal(cards, 7)
	hand.print()
	remainingDeck.print()
}
