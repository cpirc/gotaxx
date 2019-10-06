package ataxx

// Gameover ...
func (pos *Position) Gameover() bool {
	// Halfmove clock
	if pos.halfmoves >= 100 {
		return true
	}

	// No pieces left
	if pos.pieces[0].data == 0 || pos.pieces[1].data == 0 {
		return true
	}

	// No moves left
	empty := all ^ pos.pieces[0].data ^ pos.pieces[1].data ^ pos.gaps.data
	both := Bitboard{pos.pieces[0].data | pos.pieces[1].data}
	if (both.Singles().data|both.Doubles().data)&empty != 0 {
		return false
	}

	return true
}
