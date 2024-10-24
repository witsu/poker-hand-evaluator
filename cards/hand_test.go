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
			NewCard(2, 0),
			NewCard(3, 0),
			NewCard(4, 0),
			NewCard(5, 0),
			NewCard(14, 0),
		}}, "Straight Flush"},
		{Hand{cards: []Card{
			NewCard(2, 0),
			NewCard(2, 1),
			NewCard(2, 2),
			NewCard(2, 3),
			NewCard(3, 3),
		}}, "Four of a Kind"},
		{Hand{cards: []Card{
			NewCard(2, 0),
			NewCard(2, 1),
			NewCard(2, 2),
			NewCard(3, 3),
			NewCard(3, 3),
		}}, "Full House"},
		{Hand{cards: []Card{
			NewCard(2, 0),
			NewCard(3, 0),
			NewCard(4, 0),
			NewCard(13, 0),
			NewCard(14, 0),
		}}, "Flush"},
		{Hand{cards: []Card{
			NewCard(2, 0),
			NewCard(3, 0),
			NewCard(4, 0),
			NewCard(5, 2),
			NewCard(14, 1),
		}}, "Straight"},
		{Hand{cards: []Card{
			NewCard(2, 0),
			NewCard(2, 1),
			NewCard(2, 2),
			NewCard(3, 3),
			NewCard(4, 3),
		}}, "Three of a Kind"},
		{Hand{cards: []Card{
			NewCard(2, 0),
			NewCard(2, 1),
			NewCard(3, 2),
			NewCard(3, 3),
			NewCard(4, 3),
		}}, "Two Pair"},
		{Hand{cards: []Card{
			NewCard(2, 0),
			NewCard(2, 1),
			NewCard(3, 2),
			NewCard(4, 3),
			NewCard(5, 3),
		}}, "One Pair"},
		{Hand{cards: []Card{
			NewCard(2, 0),
			NewCard(3, 1),
			NewCard(4, 2),
			NewCard(7, 3),
			NewCard(8, 3),
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
