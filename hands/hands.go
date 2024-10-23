package hands

import (
	"fmt"
	"poker/cards"
	"slices"
)

func Deal(deck *cards.Deck) Hand {
	hand := Hand{
		cards: make([]cards.Card, 5),
	}
	for i := 0; i < 5; i++ {
		hand.cards[i] = deck.Deal()
	}
	slices.SortFunc(hand.cards, func(a, b cards.Card) int {
		if a.Rank < b.Rank {
			return -1
		}
		return 1
	})
	return hand
}

type rankCounts map[cards.Rank]int

type Hand struct {
	cards []cards.Card
}

func (h Hand) String() string {
	return fmt.Sprintf("%v", h.cards)
}

func (h Hand) Evaluate() string {
	ranks := make([]cards.Rank, 5)
	suits := make([]cards.Suit, 5)
	for i, card := range h.cards {
		ranks[i] = card.Rank
		suits[i] = card.Suit
	}

	isFlush := isFlush(suits)
	isStraight := isStraight(ranks)
	rankCounts := countRanks(ranks)

	switch {
	case isFlush && isStraight:
		return "Straight Flush"
	case hasNumberOfSame(rankCounts, 4):
		return "Four of a Kind"
	case hasFullHouse(rankCounts):
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

func isFlush(suits []cards.Suit) bool {
	for i := 1; i < len(suits); i++ {
		if suits[i] != suits[0] {
			return false
		}
	}
	return true
}

func isStraight(ranks []cards.Rank) bool {
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

func countRanks(ranks []cards.Rank) rankCounts {
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

func hasFullHouse(counts rankCounts) bool {
	hasThree := false
	hasTwo := false
	for _, count := range counts {
		if count == 3 {
			hasThree = true
		} else if count == 2 {
			hasTwo = true
		}
	}
	return hasThree && hasTwo
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
