package ataxx

import (
	"math/rand"
)

var SquareKeys [2][49]uint64
var TurnKey uint64

func InitZobristKeys() {
	rand.Seed(int64(143610375948384))
	for color := 0; color != 2; color++ {
		for square := 0; square != 49; square++ {
			SquareKeys[color][square] = rand.Uint64()
		}
	}
	TurnKey = rand.Uint64()
}
