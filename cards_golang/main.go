package main

func main() {
	// cards := []string{"Ace of Diamonds", newCard()}
	
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// cards := newDeck()

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
	
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	
	// cards.print()
	// hand, reaminingDeck := deal(cards, 5)
	// hand.print()
	// reaminingDeck.print()

	// fmt.Println(cards.toString())
	
	// cards.saveToFile("my_cards")
}

// func newCard() string {
// 	return "Five of Diamonds"
// }