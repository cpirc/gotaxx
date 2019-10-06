package ataxx

import (
	"fmt"
)

// Position ...
type Position struct {
	pieces    [2]Bitboard
	gaps      Bitboard
	turn      int
	halfmoves int
}

// NewPosition ...
func NewPosition(fen string) (*Position, error) {
	var position Position
	position.SetFen(fen)
	return &position, nil
}

// Turn ...
func (pos *Position) Turn() int {
	return pos.turn
}

// Set ...
func (pos *Position) Set(sq Square, piece int) {
	switch piece {
	case 0:
		pos.pieces[0].Set(sq)
		pos.pieces[1].Unset(sq)
		pos.gaps.Unset(sq)
	case 1:
		pos.pieces[0].Unset(sq)
		pos.pieces[1].Set(sq)
		pos.gaps.Unset(sq)
	case 2:
		pos.pieces[0].Unset(sq)
		pos.pieces[1].Unset(sq)
		pos.gaps.Set(sq)
	default:
		pos.pieces[0].Unset(sq)
		pos.pieces[1].Unset(sq)
		pos.gaps.Unset(sq)
	}
}

// Get ...
func (pos *Position) Get(sq Square) int {
	if pos.pieces[0].Get(sq) {
		return 0
	}
	if pos.pieces[1].Get(sq) {
		return 1
	}
	if pos.gaps.Get(sq) {
		return 2
	}
	return 3
}

// Print ...
func (pos Position) Print() {
	for i := 42; i >= 0; i++ {
		sq := Square{data: uint8(i)}
		switch pos.Get(sq) {
		case 0:
			fmt.Print("x ")
		case 1:
			fmt.Print("o ")
		case 2:
			fmt.Print("  ")
		default:
			fmt.Print("- ")
		}

		if i%7 == 6 {
			fmt.Println()
			i -= 14
		}
	}
}
