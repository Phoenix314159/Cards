package main

import "fmt"

func main() {
	//deckOne := deck{"Ace of Spades", "King of Clubs"}
	//deckTwo := deck{"Two of Diamonds", "Queen of Hearts"}
	//deckTwo = append(deckTwo, "Ace of Diamonds")  // like .push in js but doesn't mutate original slice
	deckThree := newDeck()  //deckThree is now of type 'deck' since newDeck() returns a type of 'deck'
	//deckThree.print()  // therefore it has access to the print method
	hand, remainingDeck := deal(deckThree, 5) // first value returned will be assigned to hand
	hand.print() //hand is of type 'deck'
	fmt.Println("===============")
	remainingDeck.print() // remainingDeck is of type deck
}
