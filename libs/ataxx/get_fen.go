package ataxx

import "strconv"

// GetFen ...
func (pos *Position) GetFen() string {
	var fen string = ""

	gaps := 0
	// Pieces
	for sq := 42; sq >= 0; sq++ {
		switch pos.Get(Square{uint8(sq)}) {
		case 0:
			if gaps > 0 {
				fen += strconv.Itoa(gaps)
				gaps = 0
			}
			fen += "x"
		case 1:
			if gaps > 0 {
				fen += strconv.Itoa(gaps)
				gaps = 0
			}
			fen += "o"
		case 2:
			if gaps > 0 {
				fen += strconv.Itoa(gaps)
				gaps = 0
			}
			fen += "-"
		case 3:
			gaps++
		}

		if sq%7 == 6 {
			sq -= 14
			if gaps > 0 {
				fen += strconv.Itoa(gaps)
				gaps = 0
			}
			if sq >= -1 {
				fen += "/"
			}
		}
	}

	// Turn
	if pos.turn == 0 {
		fen += " x"
	} else {
		fen += " o"
	}

	// Halfmoves
	fen += " " + strconv.Itoa(pos.halfmoves)

	return fen
}
