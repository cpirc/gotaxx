package ataxx

import (
	"errors"
	"fmt"
)

// Move ...
type Move struct {
	from Square
	to   Square
}

// NewMove ...
func NewMove(movestr string) (*Move, error) {
	if len(movestr) == 2 {
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
