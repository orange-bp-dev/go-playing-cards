package main

func main() {
	cards := newDeck()

	// for _, card := range cards {
	// 	fmt.Println(card)
	// }

	//cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

}
