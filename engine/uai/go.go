package uai

import (
	"fmt"

	"github.com/cpirc/gotaxx/engine/search"
	"github.com/cpirc/gotaxx/libs/ataxx"
)

func Go(pos ataxx.Position, input string) {
	move := search.Random(pos)
	fmt.Println("bestmove", move.String())
}
