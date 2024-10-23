package hands

import (
	"fmt"
	"poker/cards"
)

type Hand struct {
	cards []cards.Card
}

func (h Hand) String() string {
	return fmt.Sprintf("%v", h.cards)
}

func Deal(deck *cards.Deck) Hand {
	hand := Hand{
		cards: make([]cards.Card, 5),
	}
	for i := 0; i < 5; i++ {
		hand.cards[i] = deck.Deal()
	}
	return hand
}
