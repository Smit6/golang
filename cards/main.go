package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	fmt.Println()

	cards.saveToFile("cards_deck")
	newCards := newDeckFromFile("cards_deck")
	newCards.print()
	fmt.Println()

	newCards.shuffle()
	newCards.print()
}
