package main

import (
	"fmt"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newCard() deck {

	cards := deck{}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cardSuits := []string{"Spade", "Diamonds", "Hearts", "Club"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
