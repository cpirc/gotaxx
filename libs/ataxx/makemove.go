package ataxx

// Makemove ...
func (pos *Position) Makemove(move Move) {
	bbTo := Bitboard{uint64(1) << move.to.data}
	bbFrom := Bitboard{uint64(1) << move.from.data}
	neighbours := bbTo.Singles().data

	// Move our piece
	pos.pieces[pos.turn].data ^= (bbTo.data | bbFrom.data)

	// Flip captured pieces
	captured := pos.pieces[1-pos.turn].data & neighbours
	pos.pieces[pos.turn].data ^= captured
	pos.pieces[1-pos.turn].data ^= captured

	// Flip turn
	pos.turn = 1 - pos.turn

	// Halfmove counter
	pos.halfmoves++
	if captured != 0 || move.IsSingle() {
		pos.halfmoves = 0
	}
}
