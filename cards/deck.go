package main

import "fmt"

type deck []string

// function with receiver "d deck" - it acts like a member function for "deck"
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}
