package main

import "fmt"

func main() {

	//cards := newCard()
	////cards.print()
	//
	//hand, remainingDeck := deal(cards, 4)
	//hand.print()
	//remainingDeck.print()
	//
	//
	//greeting := "Hello There"
	//fmt.Println([]byte(greeting))

	cards := newCard()
	fmt.Println(cards.toString())
}
