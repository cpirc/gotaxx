package ataxx

// Move ...
type Move struct {
	from uint8
	to   uint8
}

// IsSingle ...
func (move *Move) IsSingle() bool {
	return move.from == move.to
}

// IsDouble ...
func (move *Move) IsDouble() bool {
	return move.from != move.to
}
