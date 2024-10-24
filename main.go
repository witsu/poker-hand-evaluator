package main

import (
	"fmt"
	"math/bits"
)

const (
	straightMask   uint16 = 0x001F
	aceLowStraight uint16 = 0x403c
)

type Card struct {
	rank uint16
	suit uint16
}

func evaluateHand(hand []Card) string {
	var rankBits, suitBits uint16
	var rankCount uint64
	var offset uint64

	for _, card := range hand {
		rankBits |= 1 << card.rank
		suitBits |= 1 << card.suit
		offset = 1 << (card.rank * 4)
		rankCount += offset * ((rankCount / offset & 15) + 1)
	}

	isFlush := bits.OnesCount16(suitBits) == 1
	lsb := rankBits & -rankBits
	isStraight := (rankBits/lsb == straightMask || rankBits == aceLowStraight)

	switch {
	case isStraight && isFlush:
		return "Straight Flush"
	case rankCount%15 == 1:
		return "Four of a Kind"
	case rankCount%15 == 10:
		return "Full House"
	case isFlush:
		return "Flush"
	case isStraight:
		return "Straight"
	case rankCount%15 == 9:
		return "Three of a Kind"
	case rankCount%15 == 7:
		return "Two Pair"
	case rankCount%15 == 6:
		return "One Pair"
	default:
		return "High Card"
	}
}

func main() {
	hand := []Card{
		{rank: 5, suit: 0},
		{rank: 5, suit: 1},
		{rank: 5, suit: 2},
		{rank: 7, suit: 1},
		{rank: 9, suit: 0},
	}

	result := evaluateHand(hand)
	fmt.Println("Hand evaluation:", result)
}
