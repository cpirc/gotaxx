package ataxx

// Square ...
type Square struct {
	Data uint8
}

// File ...
func (sq *Square) File() uint8 {
	return sq.Data % 7
}

// Rank ...
func (sq *Square) Rank() uint8 {
	return sq.Data / 7
}
