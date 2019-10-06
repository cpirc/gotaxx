package ataxx

// LegalMoves ...
func (pos *Position) LegalMoves(moves *[]Move) {
	legal := all ^ pos.pieces[0].data ^ pos.pieces[1].data ^ pos.gaps.data

	// Singles
	singles := Bitboard{pos.pieces[pos.turn].Singles().data & legal}
	for singles.data != 0 {
		sq := singles.LSB()
		bb := uint64(1) << sq
		singles.data ^= bb
		*moves = append(*moves, Move{sq, sq})
	}

	// Doubles
	var copy = pos.pieces[pos.turn]
	for copy.data != 0 {
		fr := copy.LSB()
		var bb = Bitboard{uint64(1) << fr}
		copy.data ^= bb.data

		doubles := Bitboard{bb.Doubles().data & legal}
		for doubles.data != 0 {
			to := doubles.LSB()
			doubles.data ^= uint64(1) << to
			*moves = append(*moves, Move{fr, to})
		}
	}
}