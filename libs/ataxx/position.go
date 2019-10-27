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
	hashKey   uint64
}

// HashPosition ...
func HashPosition(position *Position) uint64 {
	hashKey := uint64(0)
	for color, pieceBB := range position.pieces {
		for pieceBB.Data != 0 {
			square := pieceBB.LSB()
			pieceBB.Data ^= uint64(1) << square
			hashKey ^= SquareKeys[color][square]
		}
	}
	if position.turn == 1 {
		hashKey ^= TurnKey
	}
	return hashKey
}

// NewPosition ...
func NewPosition(fen string) (*Position, error) {
	var position Position
	position.SetFen(fen)
	position.hashKey = HashPosition(&position)
	return &position, nil
}

// Turn ...
func (pos *Position) Turn() int {
	return pos.turn
}

// Us ...
func (pos *Position) Us() Bitboard {
	return pos.pieces[pos.turn]
}

// Them ...
func (pos *Position) Them() Bitboard {
	return pos.pieces[1-pos.turn]
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
		sq := Square{Data: uint8(i)}
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

	fmt.Println(pos.hashKey)
}

func (pos *Position) HashKey() uint64 {
	return pos.hashKey
}
