package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	fmt.Println(hand.toString())

	hand.saveToFile("cards_on_hand")

	cards = newDeckFromFile("cards_on_hand")

	cards.print()

	cards.shuffle()

	cards.print()

}
