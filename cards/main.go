package main

import (
	"math/rand"
	"time"
)

var cards = []string{
	"Five of Diamonds",
	"Ace of Diamonds",
	"Ace of spades",
	"Six of spades",
	"Seven of spades",
	"King of hearts",
	"Queen of hearts",
	"Jack of hearts",
	"King of clubs",
	"Queen of clubs",
	"Jack of clubs",
}

func main() {
	cards := deck{randomCard(), randomCard(), randomCard()}
	// call receiver function "print()" from deck
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

func randomCard() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Int() % len(cards)
	return cards[n]
}
