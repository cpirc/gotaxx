package uai

import (
	"errors"
	"fmt"
	"strings"

	"gotaxx/libs/ataxx"
)

const STARTPOS = "x5o/7/7/7/7/7/o5x x 0"

func Position(input string) (*ataxx.Position, error) {
	parts := strings.Split(input, " ")
	if len(parts) < 2 {
		return nil, errors.New("incomplete position string")
	}

	fen, movesIdx, err := func() (string, int, error) {
		if parts[1] == "startpos" {
			return STARTPOS, 2, nil
		} else if parts[1] != "fen" {
			return "", -1, errors.New("input must start with 'position fen' or 'position startpos'")
		}

		if len(parts) < 6 {
			return "", -1, errors.New("incomplete position string")
		}
		return strings.Join(parts[2:6], " "), 6, nil
	}()
	if err != nil {
		return nil, err
	}

	pos, err := ataxx.NewPosition(fen)
	if err != nil {
		return nil, errors.New("invalid fen")
	}

	if movesIdx >= len(parts) {
		return pos, nil
	}
	if parts[movesIdx] != "moves" {
		return nil, errors.New("expected string 'moves'")
	} else if movesIdx+1 >= len(parts) {
		return nil, errors.New("expected moves after string 'moves'")
	}
	for _, moveStr := range parts[movesIdx+1:] {
		move, err := ataxx.NewMove(moveStr)
		if err != nil {
			return nil, fmt.Errorf("illegal move %s", moveStr)
		}

		pos.MakeMove(*move)
	}

	return pos, nil
}
