package ataxx

// LegalMoves ...
func (pos *Position) CaptureMoves() []Move {
	moves := make([]Move, 0, 200)

	// Can't move if it's gameover
	if pos.Gameover() == true {
		return moves
	}

	legal := pos.Them().Singles().Data ^ pos.pieces[0].Data ^ pos.pieces[1].Data ^ pos.gaps.Data

	// Singles
	singles := Bitboard{pos.pieces[pos.turn].Singles().Data & legal}
	for singles.Data != 0 {
		sq := Square{singles.LSB()}
		bb := uint64(1) << sq.Data
		singles.Data ^= bb
		moves = append(moves, Move{sq, sq})
	}

	// Doubles
	var us = pos.pieces[pos.turn]
	for us.Data != 0 {
		fr := Square{us.LSB()}
		var bb = Bitboard{uint64(1) << fr.Data}
		us.Data ^= bb.Data

		doubles := Bitboard{bb.Doubles().Data & legal}
		for doubles.Data != 0 {
			to := Square{doubles.LSB()}
			doubles.Data ^= uint64(1) << to.Data
			moves = append(moves, Move{fr, to})
		}
	}

	// Nullmove
	if len(moves) == 0 {
		moves = append(moves, NULLMOVE)
	}

	return moves
}
