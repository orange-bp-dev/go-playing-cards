package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())

	// for _, card := range cards {
	// 	fmt.Println(card)
	// }

	//cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

}
