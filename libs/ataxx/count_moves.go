package ataxx

// CountMoves ...
func (pos *Position) CountMoves() int {
	// No moves if it's gameover
	if pos.Gameover() == true {
		return 0
	}

	numMoves := 0
	legal := all ^ pos.pieces[0].data ^ pos.pieces[1].data ^ pos.gaps.data

	// Singles
	singles := Bitboard{pos.pieces[pos.turn].Singles().data & legal}
	numMoves += singles.Count()

	// Doubles
	var copy = pos.pieces[pos.turn]
	for copy.data != 0 {
		fr := copy.LSB()
		var bb = Bitboard{uint64(1) << fr}
		copy.data ^= bb.data

		doubles := Bitboard{bb.Doubles().data & legal}
		numMoves += doubles.Count()
	}

	// Nullmove
	if numMoves == 0 {
		numMoves++
	}

	return numMoves
}
