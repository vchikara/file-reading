package main

func main() {
	//var cards string = "Ace of spaces"
	cards := newDeck()
	//cards.print()
	//fmt.Println(cards)
	cards.saveToFile("my_cards")
	cards = newDeckFromFile("my_cards")
	cards.print()

}
