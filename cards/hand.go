package cards

import (
	"fmt"
)

type rankCounts map[Rank]int

type Hand struct {
	cards []Card
}

func (h Hand) Cards() []Card {
	return h.cards
}

func (h Hand) String() string {
	return fmt.Sprintf("%v", h.cards)
}

func (h Hand) Evaluate() string {
	ranks := make([]Rank, 5)
	suits := make([]Suit, 5)
	for i, card := range h.cards {
		ranks[i] = card.rank
		suits[i] = card.suit
	}

	isFlush := isFlush(suits)
	isStraight := isStraight(ranks)
	rankCounts := countRanks(ranks)

	switch {
	case isFlush && isStraight:
		return "Straight Flush"
	case hasNumberOfSame(rankCounts, 4):
		return "Four of a Kind"
	case hasNumberOfSame(rankCounts, 3) && hasNumberOfSame(rankCounts, 2):
		return "Full House"
	case isFlush:
		return "Flush"
	case isStraight:
		return "Straight"
	case hasNumberOfSame(rankCounts, 3):
		return "Three of a Kind"
	case hasTwoPair(rankCounts):
		return "Two Pair"
	case hasNumberOfSame(rankCounts, 2):
		return "One Pair"
	default:
		return "High Card"
	}
}

func isFlush(suits []Suit) bool {
	for i := 1; i < len(suits); i++ {
		if suits[i] != suits[0] {
			return false
		}
	}
	return true
}

func isStraight(ranks []Rank) bool {
	// Check for Ace-low straight (A-2-3-4-5)
	if ranks[0] == 2 && ranks[1] == 3 && ranks[2] == 4 && ranks[3] == 5 && ranks[4] == 14 {
		return true
	}
	for i := 1; i < len(ranks); i++ {
		if ranks[i] != ranks[i-1]+1 {
			return false
		}
	}
	return true
}

func countRanks(ranks []Rank) rankCounts {
	counts := make(rankCounts)
	for _, rank := range ranks {
		counts[rank]++
	}
	return counts
}

func hasNumberOfSame(counts rankCounts, n int) bool {
	for _, count := range counts {
		if count == n {
			return true
		}
	}
	return false
}

func hasTwoPair(counts rankCounts) bool {
	pairs := 0
	for _, count := range counts {
		if count == 2 {
			pairs++
		}
	}
	return pairs == 2
}
