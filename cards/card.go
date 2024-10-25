package cards

import (
	"fmt"
)

type Suit int

const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
)

func (s Suit) String() string {
	return [...]string{"Hearts", "Diamonds", "Clubs", "Spades"}[s]
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
	rank Rank
	suit Suit
}

func newCard(r Rank, s Suit) Card {
	return Card{rank: r, suit: s}
}

func (c Card) Rank() Rank {
	return c.rank
}

func (c Card) Suit() Suit {
	return c.suit
}

func (c Card) String() string {
	return fmt.Sprintf("'%s of %s'", c.rank, c.suit)
}
