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
		{"x5o/7/7/7/7/7/o5x x 0", []uint64{1, 16, 256, 6460, 155888}},
		{"x5o/7/2-1-2/7/2-1-2/7/o5x x 0", []uint64{1, 14, 196, 4184, 86528}},
		{"x5o/7/3-3/2-1-2/3-3/7/o5x x 0", []uint64{1, 16, 256, 5948, 133264}},
		{"7/7/7/7/7/7/7 x 0", []uint64{1, 0, 0, 0, 0}},
		{"7/7/7/7/ooooooo/ooooooo/xxxxxxx x 0", []uint64{1, 1, 75, 249, 14270}},
		{"7/7/7/7/ooooooo/ooooooo/xxxxxxx o 0", []uint64{1, 75, 249, 14270, 452980}},
		{"o2ox2/6o/2x4/1-3o1/-o5/-2-3/o5o o 0", []uint64{1, 68, 1532, 96695, 3017208}},
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
