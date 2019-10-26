package eval

import "gotaxx/libs/ataxx"

var Material = 1000

// @formatter:off
var Psqt = []int{
	30, 20, 10, 10, 10, 20, 30,
	20, 10, 10,  5, 10, 10, 20,
	10, 10,  5,  0,  5, 10, 10,
	10,  5,  0,  0,  0,  5, 10,
	10, 10,  5,  0,  5, 10, 10,
	20, 10, 10,  5, 10, 10, 20,
	30, 20, 10, 10, 10, 20, 30,
}
// @formatter:on

func Eval(pos *ataxx.Position) int {
	score := 0

	colors := []ataxx.Bitboard{pos.Us(), pos.Them()}
	for _, bb := range colors {
		score += bb.Count() * Material

		for bb.Data != 0 {
			sq := bb.LSB()
			bb.Data ^= uint64(1) << sq

			score += Psqt[sq]
		}

		score = -score
	}

	return score
}
