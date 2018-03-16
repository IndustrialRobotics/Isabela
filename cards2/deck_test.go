package main

import (
	"fmt"
	"testing"
)

func TestNewCard(t *testing.T) {
	decks := newCard()
	if len(decks) != 16 {
		t.Error("Test Failed - Expected 16 decks, got:", len(decks))
	} else {
		fmt.Println("Test Passed - Card number OK:", len(decks))
	}
}
