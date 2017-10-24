package main

import (
	"fmt"
	"ioutil" //implements some input output utilities functions
)

type deck []string // "extends" a base type and adds additional functionality to it

//d variable is like 'this' in js
     //receiver
func (d deck) print() { //any variable of type 'deck' get access to the print method
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck { //anytime newDeck is called it returns a type of deck
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	for _, suit := range cardSuits { // builds new deck by appending concatenated string to deck
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}
          //arguments           //return values
func deal(d deck, handSize int) (deck, deck) { //this function splits up the deck with the handSize and returns two values each of type deck
	return d[:handSize], d[handSize:] // first value returned is a 'slice' of the deck slice up to and including handSize
}	                                  // second value returned is the remaining deck

