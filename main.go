package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	// for _, card := range cards {
	// 	fmt.Println(card)
	// }

	//cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

}
