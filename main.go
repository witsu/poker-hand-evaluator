package main

import (
	"fmt"
	"poker/cards"
)

func main() {
	deck := cards.NewDeck()
	hand := deck.DealHand()
	fmt.Println(hand)
	fmt.Println(hand.Evaluate())
}
