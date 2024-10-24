package cards

import (
	"math/rand"
	"slices"
)

type Deck struct {
	cards     []Card
	dealIndex int
}

func NewDeck() *Deck {
	cards := make([]Card, 0, 52)

	for r := R2; r <= Ace; r++ {
		for s := Heart; s <= Spade; s++ {
			cards = append(cards, Card{rank: r, suit: s})
		}
	}

	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })

	return &Deck{cards: cards}
}

func (d *Deck) DealHand() Hand {
	hand := Hand{
		cards: make([]Card, 5),
	}
	for i := 0; i < 5; i++ {
		hand.cards[i] = d.dealCard()
	}
	slices.SortFunc(hand.cards, func(a, b Card) int {
		if a.rank < b.rank {
			return -1
		}
		return 1
	})
	return hand
}

func (d *Deck) dealCard() Card {
	card := d.cards[d.dealIndex]
	d.dealIndex += 1
	return card
}
