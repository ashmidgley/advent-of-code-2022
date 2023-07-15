package main

import "testing"

const TOTAL = 12

func TestGetRequiredMoveScore(t *testing.T) {
	tt := []struct {
		Name     string
		OppMove  string
		Result   int
		Expected int
	}{
		{
			Name:     "expect loss: opp plays rock",
			OppMove:  A,
			Result:   LOSS,
			Expected: SCISSORS,
		},
		{
			Name:     "expect loss: opp plays paper",
			OppMove:  B,
			Result:   LOSS,
			Expected: ROCK,
		},
		{
			Name:     "expect loss: opp plays scissors",
			OppMove:  C,
			Result:   LOSS,
			Expected: PAPER,
		},
		{
			Name:     "expect draw: opp plays rock",
			OppMove:  A,
			Result:   DRAW,
			Expected: ROCK,
		},
		{
			Name:     "expect draw: opp plays paper",
			OppMove:  B,
			Result:   DRAW,
			Expected: PAPER,
		},
		{
			Name:     "expect draw: opp plays scissors",
			OppMove:  C,
			Result:   DRAW,
			Expected: SCISSORS,
		},
		{
			Name:     "expect win: opp plays rock",
			OppMove:  A,
			Result:   WIN,
			Expected: PAPER,
		},
		{
			Name:     "expect win: opp plays paper",
			OppMove:  B,
			Result:   WIN,
			Expected: SCISSORS,
		},
		{
			Name:     "expect win: opp plays scissors",
			OppMove:  C,
			Result:   WIN,
			Expected: ROCK,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			score := getRequiredMoveScore(tc.OppMove, tc.Result)
			if score != tc.Expected {
				t.Errorf("%s - expected %d; got %d", tc.Name, tc.Expected, score)
			}
		})
	}
}

func TestCalculateTotal(t *testing.T) {
	total, _ := calculateTotal("test.txt")

	if total != TOTAL {
		t.Errorf("expected %d; got %d", TOTAL, total)
	}
}
