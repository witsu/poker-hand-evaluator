package main

import (
	"testing"
)

func Test_Evaluate(t *testing.T) {
	var tests = []struct {
		hand []Card
		want string
	}{
		{[]Card{
			{rank: 2, suit: 0},
			{rank: 3, suit: 0},
			{rank: 4, suit: 0},
			{rank: 5, suit: 0},
			{rank: 14, suit: 0},
		}, "Straight Flush"},
		{[]Card{
			{rank: 2, suit: 0},
			{rank: 2, suit: 1},
			{rank: 2, suit: 2},
			{rank: 2, suit: 3},
			{rank: 3, suit: 3},
		}, "Four of a Kind"},
		{[]Card{
			{rank: 2, suit: 0},
			{rank: 2, suit: 1},
			{rank: 2, suit: 2},
			{rank: 3, suit: 3},
			{rank: 3, suit: 3},
		}, "Full House"},
		{[]Card{
			{rank: 2, suit: 0},
			{rank: 3, suit: 0},
			{rank: 4, suit: 0},
			{rank: 13, suit: 0},
			{rank: 14, suit: 0},
		}, "Flush"},
		{[]Card{
			{rank: 2, suit: 0},
			{rank: 3, suit: 0},
			{rank: 4, suit: 0},
			{rank: 5, suit: 2},
			{rank: 14, suit: 1},
		}, "Straight"},
		{[]Card{
			{rank: 2, suit: 0},
			{rank: 2, suit: 1},
			{rank: 2, suit: 2},
			{rank: 3, suit: 3},
			{rank: 4, suit: 3},
		}, "Three of a Kind"},
		{[]Card{
			{rank: 2, suit: 0},
			{rank: 2, suit: 1},
			{rank: 3, suit: 2},
			{rank: 3, suit: 3},
			{rank: 4, suit: 3},
		}, "Two Pair"},
		{[]Card{
			{rank: 2, suit: 0},
			{rank: 2, suit: 1},
			{rank: 3, suit: 2},
			{rank: 4, suit: 3},
			{rank: 5, suit: 3},
		}, "One Pair"},
		{[]Card{
			{rank: 2, suit: 0},
			{rank: 3, suit: 1},
			{rank: 4, suit: 2},
			{rank: 7, suit: 3},
			{rank: 8, suit: 3},
		}, "High Card"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			strength := evaluateHand(tt.hand)
			if strength != tt.want {
				t.Errorf("got %s, want %s", strength, tt.want)
			}
		})
	}
}

func Benchmark_Evaluate(b *testing.B) {
	hand := []Card{
		{rank: 2, suit: 0},
		{rank: 2, suit: 1},
		{rank: 2, suit: 2},
		{rank: 3, suit: 3},
		{rank: 4, suit: 3},
	}
	for i := 0; i < b.N; i++ {
		evaluateHand(hand)
	}
}
