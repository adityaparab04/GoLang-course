package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	cards := []string{newCard(), "Queen of Hearts"}
	cards = append(cards, "Ace of Spade")
	
	for i, card := range cards {
		fmt.Println(i, card)
	}
		
}

func newCard() string {
	return "Five of Diamonds"
}