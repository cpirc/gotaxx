package ataxx

// Perft ...
func (pos Position) Perft(depth int) uint64 {
	if depth == 1 {
		return uint64(pos.CountMoves())
	} else if depth == 0 {
		return 1
	}

	var nodes uint64 = 0
	moves := make([]Move, 0, 200)
	pos.LegalMoves(&moves)

	for i := 0; i < len(moves); i++ {
		npos := pos
		npos.MakeMove(moves[i])
		nodes += npos.Perft(depth - 1)
	}

	return nodes
}
