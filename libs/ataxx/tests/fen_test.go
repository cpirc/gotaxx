package test

import (
	"testing"

	"github.com/cpirc/gotaxx/libs/ataxx"
)

// TestFen ...
func TestFen(t *testing.T) {
	fens := []string{
		"x5o/7/7/7/7/7/o5x x 0",
		"x5o/7/7/7/7/7/o5x o 0",
		"x5o/7/7/7/7/7/o5x x 50",
		"x5o/7/7/7/7/7/o5x o 50",
		"x5o/7/2-1-2/7/2-1-2/7/o5x x 0",
		"x5o/7/2-1-2/7/2-1-2/7/o5x o 0",
		"x5o/7/2-1-2/7/2-1-2/7/o5x x 100",
		"x5o/7/2-1-2/7/2-1-2/7/o5x o 100",
		"7/7/7/7/7/7/6x x 0",
		"6x/7/7/7/7/7/7 x 0",
		"x6/7/7/7/7/7/7 x 0",
		"7/7/7/7/7/7/x6 x 0",
	}

	for _, fen := range fens {
		var pos ataxx.Position
		pos.SetFen(fen)
		if fen != pos.GetFen() {
			t.Errorf("Got %s expected %s", pos.GetFen(), fen)
		}
	}
}
