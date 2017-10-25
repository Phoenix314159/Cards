package main

import "fmt"

func main() {
	//deckOne := deck{"Ace of Spades", "King of Clubs"}
	deckTwo := deck{"Two of Diamonds", "Queen of Hearts"}
	deckTwo.print()
	fmt.Println("===============")
	//deckTwo = append(deckTwo, "Ace of Diamonds")  // like .push in js but doesn't mutate original slice
	deckThree := newDeck() //deckThree is now of type 'deck' since newDeck() returns a type of 'deck'
	//deckThree.print()  // therefore it has access to the print method
	hand, remainingDeck := deal(deckThree, 5) // first value returned will be assigned to hand
	hand.print()                              //hand is of type deck
	fmt.Println("===============")
	s := remainingDeck.toString() // remainingDeck is of type deck
	fmt.Println(s)
	fmt.Println("===============")
	greeting := "Hi There!!"
	bs := []byte(greeting) //type conversion from string to a byte slice
	fmt.Println(bs)
	fmt.Println("===============")
	//deckThree.saveToFile("my_cards")
	deckFour := newDeckFromFile("my_cards")
	deckFour.print()
	fmt.Println("===============")
	deckFour.shuffle()
	deckFour.print()
	fmt.Println("===============")
}
