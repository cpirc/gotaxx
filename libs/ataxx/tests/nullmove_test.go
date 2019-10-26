package test

import (
	"testing"

	"gotaxx/libs/ataxx"
)

// TestNullmove ...
func TestNullmove(t *testing.T) {
	fens := []string{
		"7/7/7/7/ooooooo/ooooooo/xxxxxxx x 0",
		"7/7/7/7/xxxxxxx/xxxxxxx/ooooooo o 0",
	}

	movestr := ataxx.NULLMOVE.String()
	if movestr != "0000" {
		t.Errorf("Nullmove string %s", movestr)
	}

	move, err := ataxx.NewMove("0000")
	if err != nil {
		t.Errorf("Error parsing nullmove %s", err)
	} else if *move != ataxx.NULLMOVE {
		t.Errorf("NewMove(\"0000\") nullmove mismatch")
	}

	for _, fen := range fens {
		pos, _ := ataxx.NewPosition(fen)

		moves := pos.LegalMoves()

		if len(moves) != 1 {
			t.Errorf("Position %s expected 1 move, found %d moves", fen, len(moves))
		} else if moves[0] != ataxx.NULLMOVE {
			t.Errorf("Position %s expected nullmove", fen)
		}
	}
}
