package search

import (
	"math/rand"

	"github.com/cpirc/gotaxx/libs/ataxx"
)

// Random ...
func Random(pos ataxx.Position) ataxx.Move {
	var moves []ataxx.Move
	pos.LegalMoves(&moves)
	idx := rand.Intn(len(moves))
	return moves[idx]
}
