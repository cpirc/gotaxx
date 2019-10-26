package test

import (
	"testing"

	"gotaxx/engine/search"
	"gotaxx/libs/ataxx"
)

type randomTest struct {
	fen   string
	moves []string
}

// TestRandom ...
func TestRandom(t *testing.T) {
	tests := []randomTest{
		{"7/7/7/7/7/7/7 x 0", []string{"0000"}},
		{"7/7/7/7/7/7/7 o 0", []string{"0000"}},
		{"7/7/7/7/ooooooo/ooooooo/xxxxxxx x 0", []string{"0000"}},
		{"7/7/7/7/xxxxxxx/xxxxxxx/ooooooo o 0", []string{"0000"}},
		{"o6/7/7/7/7/7/6x x 0", []string{"f1", "f2", "g2", "g1e1", "g1e2", "g1e3", "g1f3", "g1g3"}},
		{"x6/7/7/7/7/7/6o o 0", []string{"f1", "f2", "g2", "g1e1", "g1e2", "g1e3", "g1f3", "g1g3"}},
	}

	for _, test := range tests {
		pos, _ := ataxx.NewPosition(test.fen)
		for i := 0; i < 50; i++ {
			move := search.Random(*pos)
			found := false

			for _, movestr := range test.moves {
				if move.String() == movestr {
					found = true
					break
				}
			}

			if found == false {
				t.Errorf("Position %s incorrect move %s", test.fen, move)
			}
		}
	}
}
