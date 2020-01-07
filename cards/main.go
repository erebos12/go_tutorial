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
}

func main() {
	cards := []string{randomCard(), randomCard(), randomCard()}
	//cards = append(cards, "Seven of spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}

func randomCard() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Int() % len(cards)
	return cards[n]
}
