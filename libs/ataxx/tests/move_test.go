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
	}

	for _, fen := range fens {
		pos := ataxx.NewPosition(fen)

		var moves []ataxx.Move
		pos.LegalMoves(&moves)

		for _, move := range moves {
			movestr := move.String()
			nmove, err := ataxx.NewMove(movestr)
			if err != nil {
				t.Errorf("Failed")
			} else if *nmove != move {
				t.Errorf("Got %s expected %s from %s", nmove, movestr, move)
			}
		}
	}
}
