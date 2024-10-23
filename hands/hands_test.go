package hands

import (
	"poker/cards"
	"testing"
)

func Test_Evaluate(t *testing.T) {
	var tests = []struct {
		hand Hand
		want string
	}{
		{Hand{cards: []cards.Card{
			cards.NewCard(2, 0),
			cards.NewCard(3, 0),
			cards.NewCard(4, 0),
			cards.NewCard(5, 0),
			cards.NewCard(14, 0),
		}}, "Straight Flush"},
		{Hand{cards: []cards.Card{
			cards.NewCard(2, 0),
			cards.NewCard(2, 1),
			cards.NewCard(2, 2),
			cards.NewCard(2, 3),
			cards.NewCard(3, 3),
		}}, "Four of a Kind"},
		{Hand{cards: []cards.Card{
			cards.NewCard(2, 0),
			cards.NewCard(2, 1),
			cards.NewCard(2, 2),
			cards.NewCard(3, 3),
			cards.NewCard(3, 3),
		}}, "Full House"},
		{Hand{cards: []cards.Card{
			cards.NewCard(2, 0),
			cards.NewCard(3, 0),
			cards.NewCard(4, 0),
			cards.NewCard(13, 0),
			cards.NewCard(14, 0),
		}}, "Flush"},
		{Hand{cards: []cards.Card{
			cards.NewCard(2, 0),
			cards.NewCard(3, 0),
			cards.NewCard(4, 0),
			cards.NewCard(5, 2),
			cards.NewCard(14, 1),
		}}, "Straight"},
		{Hand{cards: []cards.Card{
			cards.NewCard(2, 0),
			cards.NewCard(2, 1),
			cards.NewCard(2, 2),
			cards.NewCard(3, 3),
			cards.NewCard(4, 3),
		}}, "Three of a Kind"},
		{Hand{cards: []cards.Card{
			cards.NewCard(2, 0),
			cards.NewCard(2, 1),
			cards.NewCard(3, 2),
			cards.NewCard(3, 3),
			cards.NewCard(4, 3),
		}}, "Two Pair"},
		{Hand{cards: []cards.Card{
			cards.NewCard(2, 0),
			cards.NewCard(2, 1),
			cards.NewCard(3, 2),
			cards.NewCard(4, 3),
			cards.NewCard(5, 3),
		}}, "One Pair"},
		{Hand{cards: []cards.Card{
			cards.NewCard(2, 0),
			cards.NewCard(3, 1),
			cards.NewCard(4, 2),
			cards.NewCard(7, 3),
			cards.NewCard(8, 3),
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
