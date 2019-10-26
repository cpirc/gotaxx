package test

import (
	"testing"

	"gotaxx/engine/uai"
)

type positionTest struct {
	command string
	fen     string
}

// TestRandom ...
func TestUAIPosition(t *testing.T) {
	tests := []positionTest{
		{"position startpos", uai.STARTPOS},
		{"position fen 7/7/7/7/7/7/7 x 0", "7/7/7/7/7/7/7 x 0"},
		{"position fen 7/7/7/7/7/7/7 o 0", "7/7/7/7/7/7/7 o 0"},
		{"position fen 7/7/7/7/7/7/7 x 50", "7/7/7/7/7/7/7 x 50"},
		{"position fen 7/7/7/7/7/7/7 o 50", "7/7/7/7/7/7/7 o 50"},
		{"position fen x5o/7/7/7/7/7/o5x x 0 moves g2", "x5o/7/7/7/7/6x/o5x o 0"},
		{"position fen x5o/7/7/7/7/7/o5x x 0 moves g1g3", "x5o/7/7/7/6x/7/o6 o 1"},
		{"position fen x5o/7/7/7/7/7/o5x x 0 moves g2 a2", "x5o/7/7/7/7/o5x/o5x x 0"},
	}

	for _, test := range tests {
		pos, err := uai.Position(test.command)
		if err != nil {
			t.Errorf("Error parsing position command %s", test.command)
		} else if pos.GetFen() != test.fen {
			t.Errorf("Error parsing position command %s", test.command)
		}
	}
}
