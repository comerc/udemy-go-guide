package main

func main() {
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
