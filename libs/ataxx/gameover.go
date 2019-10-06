package ataxx

// Gameover ...
func (pos *Position) Gameover() bool {
	// Halfmove clock
	if pos.halfmoves >= 100 {
		return true
	}

	// No pieces left
	if pos.pieces[0].Data == 0 || pos.pieces[1].Data == 0 {
		return true
	}

	// No moves left
	empty := all ^ pos.pieces[0].Data ^ pos.pieces[1].Data ^ pos.gaps.Data
	both := Bitboard{pos.pieces[0].Data | pos.pieces[1].Data}
	if (both.Singles().Data|both.Doubles().Data)&empty != 0 {
		return false
	}

	return true
}
