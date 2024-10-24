package main

import (
	"fmt"
	"poker/cards"
)

func main() {
	deck := cards.NewDeck()
	hand := cards.DealHand(deck)
	fmt.Println(hand)
	fmt.Println(hand.Evaluate())
}
