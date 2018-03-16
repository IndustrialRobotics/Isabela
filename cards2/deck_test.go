package main

import (
	"fmt"
	"testing"
	"os"
)

func TestNewCard(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Error("FAILED: Expected 16 cards, got:", len(d))
	} else {
		fmt.Println("PASSED: Expected 16 cards, got:", len(d))
	}

	if d[len(d)-1] != "Four of Club" {
		t.Errorf("FAILED: Expected -Four of Club- and got -%v-\n", d[len(d)-1])
	} else {
		fmt.Printf("PASSED: Expected -Four of Club- and got -%v-\n", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	filename := "_deckTesting"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Error("Error", err)
	} else {
		err := os.Remove(filename)
		if err != nil {
			t.Error("Error", err)
		}

		deck := newDeck()
		deck.saveToFile(filename)
		loadedDeck := newDeckFromFile(filename)

		if len(loadedDeck) != 16 {
			t.Error("FAILED: Expected 16 and got", len(loadedDeck))
		} else {
			fmt.Println("PASSED: Expected 16 and got", len(loadedDeck))
		}

	}

}
