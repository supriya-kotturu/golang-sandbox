package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := createDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but we got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card of deck to be 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected the last card of deck to be 'Four of Clubs', but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	d := createDeck()
	d.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 in the cards deck, but got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
