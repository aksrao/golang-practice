package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("expected cards to be 16 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card as Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Diamonds" {
		t.Errorf("Expected first card as Four of Diamonds, but got %v", d[0])
	}
}
func TestSavetofile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveFile("_decktesting")
	d1 := fromFile("_decktesting")
	if len(d1) != 16 {
		t.Errorf("expected cards to be 16 but got %v", len(d))
	}
	os.Remove("_decktesting")

}
