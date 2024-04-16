package main

func main() {
	// var card string = "Ace of Spades"
	cards := newDeck()
	
	hand, remCards := deal(cards, 5)

	hand.print()
	remCards.print()
		
}