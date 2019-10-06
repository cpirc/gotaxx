package ataxx

import (
	"errors"
	"fmt"
)

// Move ...
type Move struct {
	From Square
	To   Square
}

// NULLMOVE ...
var NULLMOVE = Move{Square{49}, Square{49}}

// NewMove ...
func NewMove(movestr string) (*Move, error) {
	if movestr == "0000" {
		return &NULLMOVE, nil
	} else if len(movestr) == 2 {
		f := uint8(movestr[0] - 'a')
		r := uint8(movestr[1] - '1')
		to := Square{r*7 + f}
		return &Move{to, to}, nil
	} else if len(movestr) == 4 {
		f1 := uint8(movestr[0] - 'a')
		r1 := uint8(movestr[1] - '1')
		fr := Square{r1*7 + f1}
		f2 := uint8(movestr[2] - 'a')
		r2 := uint8(movestr[3] - '1')
		to := Square{r2*7 + f2}
		return &Move{fr, to}, nil
	}
	return nil, errors.New("Failed to parse move string")
}

// IsSingle ...
func (move *Move) IsSingle() bool {
	return move.From == move.To
}

// IsDouble ...
func (move *Move) IsDouble() bool {
	return move.From != move.To
}

func (move Move) String() string {
	if move == NULLMOVE {
		return "0000"
	} else if move.IsSingle() {
		return fmt.Sprintf("%c%c", 'a'+move.To.File(), '1'+move.To.Rank())
	}
	return fmt.Sprintf("%c%c%c%c", 'a'+move.From.File(), '1'+move.From.Rank(), 'a'+move.To.File(), '1'+move.To.Rank())
}
