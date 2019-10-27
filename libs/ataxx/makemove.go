package ataxx

import (
	"math/bits"
)

// MakeMove ...
func (pos *Position) MakeMove(move Move) {
	// Nullmove
	if move == NULLMOVE {
		pos.turn = 1 - pos.turn
		pos.hashKey ^= TurnKey
		return
	}

	bbTo := Bitboard{uint64(1) << move.To.Data}
	bbFrom := Bitboard{uint64(1) << move.From.Data}
	neighbours := bbTo.Singles().Data

	// Move our piece
	pos.pieces[pos.turn].Data ^= bbTo.Data | bbFrom.Data
	pos.hashKey ^= SquareKeys[pos.turn][move.To.Data] ^ SquareKeys[pos.turn][move.From.Data]

	// Flip captured pieces
	captured := pos.pieces[1-pos.turn].Data & neighbours
	pos.pieces[pos.turn].Data ^= captured
	pos.pieces[1-pos.turn].Data ^= captured

	// Halfmove counter
	pos.halfmoves++
	if captured != 0 || move.IsSingle() {
		pos.halfmoves = 0
	}

	// Adjust the hashKey for flipped pieces
	for captured != 0 {
		square := bits.TrailingZeros64(captured)
		captured &= captured - 1
		pos.hashKey ^= SquareKeys[1-pos.turn][square] ^ SquareKeys[pos.turn][square]
	}

	// Flip turn
	pos.turn = 1 - pos.turn
	pos.hashKey ^= TurnKey
}
