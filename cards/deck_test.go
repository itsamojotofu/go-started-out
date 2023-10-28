package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length is 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card as Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Hearts" {
		t.Errorf("Expected last card as King of Hearts, but got %v", d[len(d)-1])
	}
}
