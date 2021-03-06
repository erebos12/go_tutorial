package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var cardSuites = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
var cardValues = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Jack", "Queen", "King"}

/* Lets add custom features to []string like methods.
We create a kind of "class" here were we define all related stuff to the
'type deck []string'.
*/

type deck []string

/* General syntax for a function in Go:

func (<receiver-instance-variable> <receiver-type>)
	 nameOfFunction (varName1 typeOfVar1, ..., varNameN typeOfVarN)
	      (type-of-return-value1, ..., type-of-return-valueN)

*/

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
	for _, suit := range cardSuites { // _ == placeholder for an unused variable
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// NO RECEIVER HERE! HERE 'DECK' IS AN ARGUMENT!
// two values of type 'deck' will be returned in this function
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

// turns a 'deck' into one comma separated string
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 777)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// ds := string(bs) - conversion of byte slice (bs) into a string
	// s := strings.Split(ds, ",") - slice of string with all cards
	// return deck(s) - create a deck out of slice of string ([]string)
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() deck {
	rand.Seed(time.Now().UnixNano())
	s := []string(d)
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s
}
