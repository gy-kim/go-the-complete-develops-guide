package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")

	newCard := newDeckFromFile("my_cards")
	newCard.print()

	shuffleCards := newDeck()
	shuffleCards.shuffle()
	shuffleCards.print()
}
