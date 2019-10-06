package search

import (
	"fmt"
	"math/rand"

	"github.com/cpirc/gotaxx/libs/ataxx"
)

// Random ...
func Random(pos ataxx.Position) {
	moves := []ataxx.Move{}
	pos.LegalMoves(&moves)
	idx := rand.Intn(len(moves))
	fmt.Printf("bestmove %s\n", moves[idx])
}
