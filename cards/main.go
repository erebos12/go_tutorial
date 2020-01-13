package main

import (
	"fmt"
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
	cards := newDeck()
	hand, remainingDeck := deal(cards, 4)
	// call receiver function "print()" from deck
	fmt.Println(">>>> Your hand:")
	hand.print()
	fmt.Println(">>>> Remaining cards:")
	remainingDeck.print()
}

func randomCard() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Int() % len(cards)
	return cards[n]
}
