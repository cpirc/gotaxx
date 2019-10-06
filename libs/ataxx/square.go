package ataxx

// Square ...
type Square struct {
	data uint8
}

// File ...
func (sq *Square) File() uint8 {
	return sq.data % 7
}

// Rank ...
func (sq *Square) Rank() uint8 {
	return sq.data / 7
}
