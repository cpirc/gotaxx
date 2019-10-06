package ataxx

import (
	"strconv"
	"strings"
)

// SetFen ...
func (pos *Position) SetFen(fen string) {
	// Default
	pos.pieces[0].data = 0
	pos.pieces[1].data = 0
	pos.gaps.data = 0
	pos.turn = 0
	pos.halfmoves = 0

	results := strings.Split(fen, " ")

	// Pieces
	if len(results) >= 1 {
		var sq uint8 = 42
		for i := 0; i < len(results[0]); i++ {
			switch results[0][i] {
			case 'x':
				pos.Set(Square{sq}, 0)
				sq++
			case 'o':
				pos.Set(Square{sq}, 1)
				sq++
			case '-':
				pos.Set(Square{sq}, 2)
				sq++
			case '1':
				sq++
			case '2':
				sq += 2
			case '3':
				sq += 3
			case '4':
				sq += 4
			case '5':
				sq += 5
			case '6':
				sq += 6
			case '7':
				sq += 7
			case '/':
				sq -= 14
			}
		}
	}

	// Turn
	if len(results) >= 2 {
		if results[1] == "x" {
			pos.turn = 0
		} else {
			pos.turn = 1
		}
	}

	// Halfmove clock
	if len(results) >= 3 {
		pos.halfmoves, _ = strconv.Atoi(results[2])
	}
}
