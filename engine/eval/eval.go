package eval

import "github.com/cpirc/gotaxx/libs/ataxx"

var MATERIAL = 1000

var PST = []int{
	30, 20, 10, 10, 10, 20, 30, 20, 10, 10, 5, 10, 10,
	20, 10, 10, 5, 0, 5, 10, 10, 10, 5, 0, 0, 0,
	5, 10, 10, 10, 5, 0, 5, 10, 10, 20, 10, 10, 5,
	10, 10, 20, 30, 20, 10, 10, 10, 20, 30,
}

func Eval(pos *ataxx.Position) int {
	score := 0

	colors := []ataxx.Bitboard{pos.Us(), pos.Them()}
	for _, bb := range colors {
		score += bb.Count() * MATERIAL

		for bb.Data != 0 {
			sq := bb.LSB()
			bb.Data ^= uint64(1) << sq

			score += PST[sq]
		}

		score = -score
	}

	return score
}
