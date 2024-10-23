package main

import (
	"fmt"
	"poker/cards"
	"poker/hands"
)

func main() {
	deck := cards.NewDeck()
	hand := hands.Deal(deck)

	fmt.Println(hand)
}
