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
			newCard(R2, Heart),
			newCard(R3, Heart),
			newCard(R4, Heart),
			newCard(R5, Heart),
			newCard(Ace, Heart),
		}}, "Straight Flush"},
		{Hand{cards: []Card{
			newCard(R2, Heart),
			newCard(R2, Diamond),
			newCard(R2, Clubs),
			newCard(R2, Spade),
			newCard(R3, Spade),
		}}, "Four of a Kind"},
		{Hand{cards: []Card{
			newCard(R2, Heart),
			newCard(R2, Diamond),
			newCard(R2, Clubs),
			newCard(R3, Spade),
			newCard(R3, Spade),
		}}, "Full House"},
		{Hand{cards: []Card{
			newCard(R2, Heart),
			newCard(R3, Heart),
			newCard(R4, Heart),
			newCard(King, Heart),
			newCard(Ace, Heart),
		}}, "Flush"},
		{Hand{cards: []Card{
			newCard(R2, Heart),
			newCard(R3, Heart),
			newCard(R4, Heart),
			newCard(R5, Clubs),
			newCard(Ace, Diamond),
		}}, "Straight"},
		{Hand{cards: []Card{
			newCard(R2, Heart),
			newCard(R2, Diamond),
			newCard(R2, Clubs),
			newCard(R3, Spade),
			newCard(R4, Spade),
		}}, "Three of a Kind"},
		{Hand{cards: []Card{
			newCard(R2, Heart),
			newCard(R2, Diamond),
			newCard(R3, Clubs),
			newCard(R3, Spade),
			newCard(R4, Spade),
		}}, "Two Pair"},
		{Hand{cards: []Card{
			newCard(R2, Heart),
			newCard(R2, Diamond),
			newCard(R3, Clubs),
			newCard(R4, Spade),
			newCard(R5, Spade),
		}}, "One Pair"},
		{Hand{cards: []Card{
			newCard(R2, Heart),
			newCard(R3, Diamond),
			newCard(R4, Clubs),
			newCard(R7, Spade),
			newCard(R8, Spade),
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
