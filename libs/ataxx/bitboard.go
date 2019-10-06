package ataxx

import "math/bits"

const (
	all uint64 = 0x1FFFFFFFFFFFF
	// Files
	fileA uint64 = 0x0040810204081
	fileB uint64 = 0x0081020408102
	fileF uint64 = 0x0810204081020
	fileG uint64 = 0x1020408102040
	// Ranks
	// Not Files
	notFileA uint64 = 0x1fbf7efdfbf7e
	notFileB uint64 = 0x1f7efdfbf7efd
	notFileF uint64 = 0x17efdfbf7efdf
	notFileG uint64 = 0x0fdfbf7efdfbf
	// Stuff
	fileAB    uint64 = fileA | fileB
	fileFG    uint64 = fileF | fileG
	notFileAB uint64 = notFileA & notFileB
	notFileFG uint64 = notFileF & notFileG
)

// Bitboard ...
type Bitboard struct {
	data uint64
}

// Set ...
func (bb *Bitboard) Set(sq Square) {
	bb.data |= 1 << sq.data
}

// Unset ...
func (bb *Bitboard) Unset(sq Square) {
	bb.data &= ^(1 << sq.data)
}

// Get ...
func (bb *Bitboard) Get(sq Square) bool {
	return bb.data&(1<<sq.data) != 0
}

// Count ...
func (bb *Bitboard) Count() int {
	return bits.OnesCount64(bb.data)
}

// LSB ...
func (bb *Bitboard) LSB() uint8 {
	return uint8(bits.TrailingZeros64(bb.data))
}

// North ...
func (bb *Bitboard) North() Bitboard {
	return Bitboard{(bb.data << 7) & all}
}

// South ...
func (bb *Bitboard) South() Bitboard {
	return Bitboard{bb.data >> 7}
}

// East ...
func (bb *Bitboard) East() Bitboard {
	return Bitboard{(bb.data << 1) & notFileA}
}

// West ...
func (bb *Bitboard) West() Bitboard {
	return Bitboard{(bb.data >> 1) & notFileG}
}

// NorthEast ...
func (bb *Bitboard) NorthEast() Bitboard {
	return Bitboard{(bb.data << 8) & notFileA}
}

// NorthWest ...
func (bb *Bitboard) NorthWest() Bitboard {
	return Bitboard{(bb.data << 6) & notFileG}
}

// SouthEast ...
func (bb *Bitboard) SouthEast() Bitboard {
	return Bitboard{(bb.data >> 6) & notFileA}
}

// SouthWest ...
func (bb *Bitboard) SouthWest() Bitboard {
	return Bitboard{(bb.data >> 8) & notFileG}
}

// Singles ...
func (bb *Bitboard) Singles() Bitboard {
	return Bitboard{
		bb.NorthEast().data |
			bb.NorthWest().data |
			bb.SouthEast().data |
			bb.SouthWest().data |
			bb.North().data |
			bb.South().data |
			bb.East().data |
			bb.West().data}
}

// Doubles ...
func (bb *Bitboard) Doubles() Bitboard {
	var moves uint64 = 0
	var asd = bb.data
	moves |= (asd << 12) & notFileFG // North North West West
	moves |= (asd << 13) & notFileG  // North North West
	moves |= (asd << 14)             // North North
	moves |= (asd << 15) & notFileA  // North North East
	moves |= (asd << 16) & notFileAB // North North East East

	moves |= (asd >> 16) & notFileFG // South South West West
	moves |= (asd >> 15) & notFileG  // South South West
	moves |= (asd >> 14)             // South South
	moves |= (asd >> 13) & notFileA  // South South East
	moves |= (asd >> 12) & notFileAB // South South East East

	moves |= (asd << 9) & notFileAB // East East North
	moves |= (asd << 2) & notFileAB // East East
	moves |= (asd >> 5) & notFileAB // East East South

	moves |= (asd << 5) & notFileFG // West West North
	moves |= (asd >> 2) & notFileFG // West West
	moves |= (asd >> 9) & notFileFG // West West South

	return Bitboard{moves}
}
