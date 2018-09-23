package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck of length 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs but got %v", d[len(d)-1])
	}

}

func TestNewDeckandLoadFromFile(t *testing.T) {
	//remove files from last test, if any are still not deleted. Clean up
	os.Remove("_testing")

	d := newDeck()

	d.saveToFile("_testing")

	loadedDeck := newDeckFromFile("_testing")

	if len(d) != len(loadedDeck) {
		t.Errorf("Both decks are not same, original %v, got %v", len(d), len(loadedDeck))
	}

	os.Remove("_testing")
}
