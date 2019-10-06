package ataxx

// LegalMoves ...
func (pos *Position) LegalMoves(moves *[]Move) {
	// Can't move if it's gameover
	if pos.Gameover() == true {
		return
	}

	legal := all ^ pos.pieces[0].data ^ pos.pieces[1].data ^ pos.gaps.data

	// Singles
	singles := Bitboard{pos.pieces[pos.turn].Singles().data & legal}
	for singles.data != 0 {
		sq := Square{singles.LSB()}
		bb := uint64(1) << sq.data
		singles.data ^= bb
		*moves = append(*moves, Move{sq, sq})
	}

	// Doubles
	var copy = pos.pieces[pos.turn]
	for copy.data != 0 {
		fr := Square{copy.LSB()}
		var bb = Bitboard{uint64(1) << fr.data}
		copy.data ^= bb.data

		doubles := Bitboard{bb.Doubles().data & legal}
		for doubles.data != 0 {
			to := Square{doubles.LSB()}
			doubles.data ^= uint64(1) << to.data
			*moves = append(*moves, Move{fr, to})
		}
	}

	// Nullmove
	if len(*moves) == 0 {
		*moves = append(*moves, NULLMOVE)
	}
}
