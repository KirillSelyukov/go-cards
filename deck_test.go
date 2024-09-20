package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	expectedLen := 52

	if len(deck) != expectedLen {
		t.Errorf("expected deck length of %v but got %v", expectedLen, len(deck))
	}
	expectedFirstCard := "Ace of Spades"

	if deck[0] != expectedFirstCard {
		t.Errorf("Expect first card to be %v but got %v", expectedFirstCard, deck[0])
	}
	expectedLastCard := "Two of Clubs"

	if deck[len(deck)-1] != expectedLastCard {
		t.Errorf("Expect last card to be %v but got %v", expectedLastCard, deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)
	expectedLen := 52
	if len(loadedDeck) != expectedLen {
		t.Errorf("Expected %v cards in deck, got %v", expectedLen, len(loadedDeck))
	}

	os.Remove(filename)
}
