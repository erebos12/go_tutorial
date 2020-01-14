package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	el := 36
	if len(d) != el {
		t.Errorf("Expected deck len of %v, but got %v", el, len(d))
	}
	es := "Ace of Spades"
	if d[0] != es {
		t.Errorf("Expected deck len of %v, but got %v", es, d[0])
	}
	es = "King of Clubs"
	if d[el-1] != es {
		t.Errorf("Expected deck len of %v, but got %v", es, d[el-1])
	}
}
