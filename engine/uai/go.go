package uai

import (
	"errors"
	"fmt"

	"github.com/cpirc/gotaxx/engine/options"
	"github.com/cpirc/gotaxx/engine/search"
	"github.com/cpirc/gotaxx/libs/ataxx"
)

func Go(pos ataxx.Position, input string, stop chan struct{}) {
	searchType, _ := options.GetCombo("SearchType")

	move, err := func() (ataxx.Move, error) {
		if searchType == "AlphaBeta" {
			return search.AlphaBeta(pos, stop), nil
		} else if searchType == "MostCaptures" {
			return search.MostCaptures(pos), nil
		} else if searchType == "Random" {
			return search.Random(pos), nil
		} else {
			return ataxx.Move{}, errors.New("could not find SearchType " + searchType)
		}
	}()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("bestmove", move.String())
}
