package search

import (
	"math/rand"

	"gotaxx/libs/ataxx"
)

// MostCaptures ...
func MostCaptures(pos ataxx.Position) ataxx.Move {
	moves := pos.LegalMoves()

	if len(moves) <= 0 || moves[0] == ataxx.NULLMOVE {
		return ataxx.NULLMOVE
	}

	var bestMoves []ataxx.Move
	bestScore := -1

	for _, move := range moves {
		bb := ataxx.Bitboard{Data: 1 << move.To.Data}
		captured := ataxx.Bitboard{Data: bb.Singles().Data & pos.Them().Data}
		score := captured.Count()

		if move.IsSingle() == true {
			score++
		}

		if score > bestScore {
			bestScore = score
			bestMoves = bestMoves[:0]
			bestMoves = append(bestMoves, move)
		} else if score == bestScore {
			bestMoves = append(bestMoves, move)
		}
	}

	idx := rand.Intn(len(bestMoves))
	return bestMoves[idx]
}
