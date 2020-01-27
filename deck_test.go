package main

import (
	"os"
	"testing"
)

// uppercase first character in test function names.
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected 52 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards but got %v", len(loadedDeck))
	}
}
