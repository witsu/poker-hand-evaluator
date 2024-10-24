package main

import (
	"fmt"
	"poker/cards"
)

func main() {
	deck := cards.NewDeck()
	hand := cards.Deal(deck)
	fmt.Println(hand)
	fmt.Println(hand.Evaluate())
}
