package test

import (
	"strconv"
	"testing"

	"gotaxx/libs/ataxx"
)

type gameoverTest struct {
	fen      string
	gameover bool
}

// TestGameover ...
func TestGameover(t *testing.T) {
	tests := []gameoverTest{
		{"x5o/7/7/7/7/7/o5x x 0", false},
		{"x5o/7/2-1-2/7/2-1-2/7/o5x x 0", false},
		{"x5o/7/3-3/2-1-2/3-3/7/o5x x 0", false},
		{"7/7/7/7/7/7/7 x 0", true},
		{"7/7/7/7/7/7/7 o 0", true},
		{"x5o/7/7/7/7/7/o5x x 0", false},
		{"x5o/7/7/7/7/7/o5x x 99", false},
		{"x5o/7/7/7/7/7/o5x x 100", true},
		{"x5o/7/7/7/7/7/o5x x 101", true},
		{"x6/7/7/7/7/7/7 x 0", true},
		{"x6/7/7/7/7/7/7 o 0", true},
		{"o6/7/7/7/7/7/7 x 0", true},
		{"o6/7/7/7/7/7/7 o 0", true},
		{"7/7/7/7/7/-------/oooxxxx x 0", false},
		{"7/7/7/7/-------/-------/oooxxxx x 0", true},
		{"oooxxxx/oooxxxx/oooxxxx/ooo1xxx/oooxxxx/oooxxxx/oooxxxx x 0", false},
		{"oooxxxx/oooxxxx/oooxxxx/ooo1xxx/oooxxxx/oooxxxx/oooxxxx o 0", false},
		{"oooxxxx/oooxxxx/oooxxxx/oooxxxx/oooxxxx/oooxxxx/oooxxxx x 0", true},
		{"oooxxxx/oooxxxx/oooxxxx/oooxxxx/oooxxxx/oooxxxx/oooxxxx o 0", true},
	}

	for _, test := range tests {
		pos, _ := ataxx.NewPosition(test.fen)
		if pos.Gameover() != test.gameover {
			t.Errorf("Position %s expected %s", test.fen, strconv.FormatBool(test.gameover))
		}
	}
}
