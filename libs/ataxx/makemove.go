package ataxx

// MakeMove ...
func (pos *Position) MakeMove(move Move) {
	// Nullmove
	if move == NULLMOVE {
		pos.turn = 1 - pos.turn
		return
	}

	bbTo := Bitboard{uint64(1) << move.To.Data}
	bbFrom := Bitboard{uint64(1) << move.From.Data}
	neighbours := bbTo.Singles().Data

	// Move our piece
	pos.pieces[pos.turn].Data ^= (bbTo.Data | bbFrom.Data)

	// Flip captured pieces
	captured := pos.pieces[1-pos.turn].Data & neighbours
	pos.pieces[pos.turn].Data ^= captured
	pos.pieces[1-pos.turn].Data ^= captured

	// Flip turn
	pos.turn = 1 - pos.turn

	// Halfmove counter
	pos.halfmoves++
	if captured != 0 || move.IsSingle() {
		pos.halfmoves = 0
	}
}
