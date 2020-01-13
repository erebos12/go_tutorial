package main

import "fmt"

/* Lets add custom features to []string like methods.
We create a kind of "class" here were we define all related stuff to the
'type deck []string'.
*/

type deck []string

// function with receiver "d deck" - it acts like a member function for "deck"
// receiver on a function
func (d deck) print() {
	for i, card := range d { // d is similar to this or self
		fmt.Println(i, card)
	}
}

// similar to constructor in OO universe
//returns a value of type deck
func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Jack", "Quenn", "King"}
	// _ == placeholder for an unused variable
	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
