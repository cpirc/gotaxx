package ataxx

import "fmt"

// Move ...
type Move struct {
	from Square
	to   Square
}

// IsSingle ...
func (move *Move) IsSingle() bool {
	return move.from == move.to
}

// IsDouble ...
func (move *Move) IsDouble() bool {
	return move.from != move.to
}
