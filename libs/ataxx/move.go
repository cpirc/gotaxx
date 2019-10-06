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

func (move Move) String() string {
	if move.IsSingle() {
		return fmt.Sprintf("%c%c", 'a'+move.to.File(), '1'+move.to.Rank())
	}
	return fmt.Sprintf("%c%c%c%c", 'a'+move.from.File(), '1'+move.from.Rank(), 'a'+move.to.File(), '1'+move.to.Rank())
}
