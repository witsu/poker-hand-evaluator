package cards

import (
	"fmt"

	"math/rand"
)

type Suit int

const (
	Heart Suit = iota
	Diamond
	Clubs
	Spade
)

func (s Suit) String() string {
	return [...]string{"Heart", "Diamond", "Clubs", "Spade"}[s]
}

type Rank int

const (
	R2 Rank = iota + 2
	R3
	R4
	R5
	R6
	R7
	R8
	R9
	R10
	Jack
	Queen
	King
	Ace
)

func (r Rank) String() string {
	return [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}[r-2]
}

type Card struct {
	Rank Rank
	Suit Suit
}

func NewCard(r, s int) Card {
	return Card{Rank: Rank(r), Suit: Suit(s)}
}

func (c Card) String() string {
	return fmt.Sprintf("'%s of %s'", c.Rank, c.Suit)
}

type Deck struct {
	cards     []Card
	dealIndex int
}

func NewDeck() *Deck {
	cards := make([]Card, 0, 52)

	for r := R2; r <= Ace; r++ {
		for s := Heart; s <= Spade; s++ {
			cards = append(cards, Card{Rank: r, Suit: s})
		}
	}

	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })

	return &Deck{cards: cards}
}

func (d *Deck) Deal() Card {
	card := d.cards[d.dealIndex]
	d.dealIndex += 1
	return card
}
