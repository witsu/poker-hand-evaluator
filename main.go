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

// Evaluate poker 5 card hand using bits to represent ranks and suits
// Algorithm from https://jonathanhsiao.com/blog/evaluating-poker-hands-with-bit-math
func evaluateHand(hand []Card) string {
	// rankBits stores info which rank was used by using bit per rank [A K Q J 10 9 8 7 6 5 4 3 2 _ _]
	// example hand [K, K, Q, J, 10] is stored as [011110000000000]
	var rankBits uint16

	//  suitBits stores info which suit was used
	// example [Hearts, Hearts, Diamond, Diamond, Diamond] [0011]
	var suitBits uint16

	// rankCount stores groups of 4 bits per rank to count ranks
	// example hand [K, K, Q, J, 10] is stored as [0000 0011 0001 0001 0001 0000 0000...0000]
	var rankCount, rankCountOffset uint64

	for _, card := range hand {
		rankBits |= 1 << card.rank
		suitBits |= 1 << card.suit
		rankCountOffset = 1 << (card.rank * 4)
		rankCount += rankCountOffset * ((rankCount / rankCountOffset & 15) + 1)
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
