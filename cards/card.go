package cards

import (
	"fmt"
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
	rank Rank
	suit Suit
}

func newCard(r Rank, s Suit) Card {
	return Card{rank: r, suit: s}
}

func (c Card) String() string {
	return fmt.Sprintf("'%s of %s'", c.rank, c.suit)
}
