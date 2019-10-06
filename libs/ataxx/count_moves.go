package ataxx

// CountMoves ...
func (pos *Position) CountMoves() int {
	// No moves if it's gameover
	if pos.Gameover() == true {
		return 0
	}

	numMoves := 0
	legal := all ^ pos.pieces[0].Data ^ pos.pieces[1].Data ^ pos.gaps.Data

	// Singles
	singles := Bitboard{pos.pieces[pos.turn].Singles().Data & legal}
	numMoves += singles.Count()

	// Doubles
	var copy = pos.pieces[pos.turn]
	for copy.Data != 0 {
		fr := copy.LSB()
		var bb = Bitboard{uint64(1) << fr}
		copy.Data ^= bb.Data

		doubles := Bitboard{bb.Doubles().Data & legal}
		numMoves += doubles.Count()
	}

	// Nullmove
	if numMoves == 0 {
		numMoves++
	}

	return numMoves
}
