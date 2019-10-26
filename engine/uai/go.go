package uai

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gotaxx/engine/options"
	"gotaxx/engine/search"
	"gotaxx/libs/ataxx"
)

func Go(pos ataxx.Position, input string, stop chan struct{}, searchStopper func()) {
	movetime := time.Duration(0)

	parts := strings.Split(input, " ")
	for i := 1; i < len(parts); i += 2 {
		part := parts[i]
		if part == "movetime" {
			movetimeMillis, err := strconv.Atoi(parts[i+1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			movetime = time.Duration(movetimeMillis) * time.Millisecond
		}
	}

	if movetime > time.Duration(0) {
		go func() {
			time.Sleep(movetime)
			searchStopper()
		}()
	}

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
