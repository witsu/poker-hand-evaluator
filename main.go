package main

import (
	"fmt"
	"poker/cards"
)

func main() {
	deck := cards.NewDeck()
	hand := make([]cards.Card, 5)
	for i := 0; i < 5; i++ {
		hand[i] = deck.Deal()
	}
	for _, c := range hand {
		fmt.Println(c)
	}
}
