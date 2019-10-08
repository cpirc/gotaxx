package test

import (
	"testing"

	"github.com/cpirc/gotaxx/libs/ataxx"
)

// TestMove ...
func TestMove(t *testing.T) {
	fens := []string{
		"x5o/7/7/7/7/7/o5x x 0",
		"x5o/7/2-1-2/7/2-1-2/7/o5x x 0",
		"x5o/7/7/7/7/7/o5x x 99",
		"x5o/7/7/7/7/7/o5x x 100",
		"x5o/7/7/7/7/7/o5x x 101",
	}

	for _, fen := range fens {
		pos, _ := ataxx.NewPosition(fen)

		moves := pos.LegalMoves()

		// CountMove
		if pos.CountMoves() != len(moves) {
			t.Errorf("Position %s CountMoves mismatch %d %d", fen, pos.CountMoves(), len(moves))
		}

		// Move to string
		for _, move := range moves {
			movestr := move.String()
			nmove, err := ataxx.NewMove(movestr)
			if err != nil {
				t.Errorf("Failed to parse movestr %s", movestr)
			} else if *nmove != move {
				t.Errorf("Got %s expected %s", nmove, movestr)
			}
		}
	}
}
