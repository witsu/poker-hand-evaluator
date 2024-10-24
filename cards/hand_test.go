package cards

import (
	"testing"
)

func Test_Evaluate(t *testing.T) {
	var tests = []struct {
		hand Hand
		want string
	}{
		{Hand{cards: []Card{
			newCard(R2, Hearts),
			newCard(R3, Hearts),
			newCard(R4, Hearts),
			newCard(R5, Hearts),
			newCard(Ace, Hearts),
		}}, "Straight Flush"},
		{Hand{cards: []Card{
			newCard(R2, Hearts),
			newCard(R2, Diamonds),
			newCard(R2, Clubs),
			newCard(R2, Spades),
			newCard(R3, Spades),
		}}, "Four of a Kind"},
		{Hand{cards: []Card{
			newCard(R2, Hearts),
			newCard(R2, Diamonds),
			newCard(R2, Clubs),
			newCard(R3, Spades),
			newCard(R3, Spades),
		}}, "Full House"},
		{Hand{cards: []Card{
			newCard(R2, Hearts),
			newCard(R3, Hearts),
			newCard(R4, Hearts),
			newCard(King, Hearts),
			newCard(Ace, Hearts),
		}}, "Flush"},
		{Hand{cards: []Card{
			newCard(R2, Hearts),
			newCard(R3, Hearts),
			newCard(R4, Hearts),
			newCard(R5, Clubs),
			newCard(Ace, Diamonds),
		}}, "Straight"},
		{Hand{cards: []Card{
			newCard(R2, Hearts),
			newCard(R2, Diamonds),
			newCard(R2, Clubs),
			newCard(R3, Spades),
			newCard(R4, Spades),
		}}, "Three of a Kind"},
		{Hand{cards: []Card{
			newCard(R2, Hearts),
			newCard(R2, Diamonds),
			newCard(R3, Clubs),
			newCard(R3, Spades),
			newCard(R4, Spades),
		}}, "Two Pair"},
		{Hand{cards: []Card{
			newCard(R2, Hearts),
			newCard(R2, Diamonds),
			newCard(R3, Clubs),
			newCard(R4, Spades),
			newCard(R5, Spades),
		}}, "One Pair"},
		{Hand{cards: []Card{
			newCard(R2, Hearts),
			newCard(R3, Diamonds),
			newCard(R4, Clubs),
			newCard(R7, Spades),
			newCard(R8, Spades),
		}}, "High Card"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			strength := tt.hand.Evaluate()
			if strength != tt.want {
				t.Errorf("got %s, want %s", strength, tt.want)
			}
		})
	}
}
