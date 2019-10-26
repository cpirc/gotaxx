package test

import (
	"testing"

	"gotaxx/engine/search"
	"gotaxx/libs/ataxx"
)

type mostCapturesTest struct {
	fen   string
	moves []string
}

// TestMostCaptures ...
func TestMostCaptures(t *testing.T) {
	tests := []mostCapturesTest{
		{"7/7/7/7/7/7/7 x 0", []string{"0000"}},
		{"7/7/7/7/7/7/7 o 0", []string{"0000"}},
		{"7/7/7/7/ooooooo/ooooooo/xxxxxxx x 0", []string{"0000"}},
		{"7/7/7/7/xxxxxxx/xxxxxxx/ooooooo o 0", []string{"0000"}},
		{"7/7/7/7/xxxxxxx/xxxxxxx/ooooooo x 0", []string{"a4", "b4", "c4", "d4", "e4", "f4", "g4"}},
		{"7/7/7/7/ooooooo/ooooooo/xxxxxxx o 0", []string{"a4", "b4", "c4", "d4", "e4", "f4", "g4"}},
		{"x5o/7/7/7/7/7/o5x x 0", []string{"f1", "f2", "g2", "a6", "b6", "b7"}},
		{"x5o/7/7/7/7/7/o5x o 0", []string{"b1", "b2", "a2", "g6", "f6", "f7"}},
		{"x5o/7/7/7/7/7/3o2x x 0", []string{"g1e1", "g1e2", "f1", "f2", "g2", "a6", "b6", "b7"}},
		{"x5o/7/7/7/7/3o3/3o2x x 0", []string{"g1e1", "g1e2"}},
	}

	for _, test := range tests {
		pos, _ := ataxx.NewPosition(test.fen)
		for i := 0; i < 50; i++ {
			move := search.MostCaptures(*pos)
			found := false

			for _, movestr := range test.moves {
				if move.String() == movestr {
					found = true
					break
				}
			}

			if found == false {
				t.Errorf("Position %s wrong move %s", test.fen, move)
			}
		}
	}
}
