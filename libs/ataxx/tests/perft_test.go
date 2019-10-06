package test

import (
	"testing"

	"github.com/cpirc/gotaxx/libs/ataxx"
)

type test struct {
	fen   string
	nodes []uint64
}

// TestPerft ...
func TestPerft(t *testing.T) {
	tests := []test{
		{"x5o/7/7/7/7/7/o5x x 0", []uint64{1, 16, 256, 6460}},
		{"x5o/7/2-1-2/7/2-1-2/7/o5x x 0", []uint64{1, 14, 196, 4184}},
	}

	for _, test := range tests {
		var pos ataxx.Position
		pos.SetFen(test.fen)
		for i := 0; i < len(test.nodes); i++ {
			n := pos.Perft(i)
			if n != test.nodes[i] {
				t.Errorf("Fen %s Got %d expected %d", test.fen, n, test.nodes[i])
			}
		}
	}
}
