package search

import (
	"fmt"
	"time"

	"gotaxx/engine/eval"
	"gotaxx/libs/ataxx"
)

type Result struct {
	score int
	pv    []ataxx.Move
}

func AlphaBeta(pos ataxx.Position, stop chan struct{}) ataxx.Move {
	startTime := time.Now()
	var bestMove ataxx.Move

IdLoop:
	for depth := 1; depth < 128; depth++ {
		latestResult := AlphaBetaImpl(pos, -30000, +30000, depth, 0, stop)
		timeTaken := time.Since(startTime)

		if depth > 1 {
			select {
			case <-stop:
				break IdLoop
			default:
			}
		}

		bestMove = latestResult.pv[0]
		fmt.Printf("info score %d depth %d time %s moves", latestResult.score, depth, timeTaken)
		for _, move := range latestResult.pv {
			fmt.Print(" ", move.String())
		}
		fmt.Println()
	}

	return bestMove
}

func AlphaBetaImpl(pos ataxx.Position, alpha int, beta int, depth int, ply int, stop chan struct{}) Result {
	if depth <= 0 {
		return Result{eval.Eval(&pos), nil}
	}

	if ply > 0 {
		select {
		case <-stop:
			return Result{eval.Eval(&pos), nil}
		default:
		}

		if ply >= 128 {
			return Result{eval.Eval(&pos), nil}
		}
	}

	moves := pos.LegalMoves()

	if len(moves) <= 0 || moves[0] == ataxx.NULLMOVE {
		return Result{eval.Eval(&pos), []ataxx.Move{ataxx.NULLMOVE}}
	}

	var pv []ataxx.Move
	bestScore := -30000
	for _, move := range moves {
		childPos := pos
		childPos.MakeMove(move)

		result := AlphaBetaImpl(childPos, -beta, -alpha, depth-1, ply+1, stop)
		result.score = -result.score

		if ply > 0 {
			select {
			case <-stop:
				return Result{eval.Eval(&pos), nil}
			default:
			}
		}

		if result.score > bestScore {
			bestScore = result.score
		}
		if bestScore > alpha {
			alpha = bestScore

			pv = []ataxx.Move{}
			pv = append(pv, move)
			if len(result.pv) > 0 {
				pv = append(pv, result.pv...)
			}
		}
		if bestScore >= beta {
			break
		}
	}

	return Result{bestScore, pv}
}
