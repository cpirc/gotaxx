package ataxx

import (
	"gotaxx/engine/perft_tt"
)

// Perft ...
func (pos Position) Perft(depth int) uint64 {
	if depth == 1 {
		return uint64(pos.CountMoves())
	} else if depth == 0 {
		return 1
	}

	var nodes uint64 = 0
	moves := pos.LegalMoves()

	for i := 0; i < len(moves); i++ {
		npos := pos
		npos.MakeMove(moves[i])
		nodes += npos.Perft(depth - 1)
	}

	return nodes
}

// Perft ...
func (pos Position) HashPerft(depth int) uint64 {
	if depth == 1 {
		return uint64(pos.CountMoves())
	} else if depth == 0 {
		return 1
	}

	ttEntry := perft_tt.PerftTranspositionTable.Probe(pos.hashKey)
	if ttEntry.Key == pos.hashKey && ttEntry.Depth == depth {
		return ttEntry.NodeCount
	}

	var nodes uint64 = 0
	moves := pos.LegalMoves()
	for i := 0; i < len(moves); i++ {
		npos := pos
		npos.MakeMove(moves[i])
		nodes += npos.Perft(depth - 1)
	}

	perft_tt.PerftTranspositionTable.Insert(perft_tt.PerftEntry{
		Depth:     depth,
		NodeCount: nodes,
		Key:       pos.hashKey,
	})

	return nodes
}
