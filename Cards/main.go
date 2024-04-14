package main

func main() {
	// var card string = "Ace of Spades"
	cards := deck{newCard(), "Queen of Hearts"}
	cards = append(cards, "Ace of Spade")
	
	cards.print()
		
}

func newCard() string {
	return "Five of Diamonds"
}