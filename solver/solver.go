package solver

import (
	connectfour "github.com/jacobomantilla10/connect-four"
)

var moveOrder = [7]int{3, 2, 4, 1, 5, 0, 6}

func GetBestMove(b connectfour.Board) int {
	bestMove := 0
	bestScore := 1000

	for _, move := range moveOrder {
		if b.CanPlay(move) && b.IsWinningMove(move) {
			return move
		}
	}

	for _, move := range moveOrder {
		newBoard := b

		if !newBoard.CanPlay(move) {
			continue
		}

		newBoard.Play(move)
		branchScore := Negamax(newBoard, -1000, 1000)
		if branchScore < bestScore {
			bestMove = move
			bestScore = branchScore
		}
	}
	return bestMove
}

func Negamax(b connectfour.Board, alpha, beta int) int {
	alphaOrig := alpha

	// TODO query transpositiontable to get value and make check the flag
	tt := Get(b.Key())
	if tt.Val != -999 {
		if tt.Flag == EXACT {
			return tt.Val
		} else if tt.Flag == LOWER {
			alpha = max(alpha, tt.Val)
		} else if tt.Flag == UPPER {
			beta = min(beta, tt.Val)
		}

		if alpha >= beta {
			return tt.Val
		}
	}

	if b.IsDrawn() {
		return 0
	}

	for _, move := range moveOrder {
		if b.CanPlay(move) && b.IsWinningMove(move) {
			return ((7*6 + 1 - b.NumMoves()) / 2)
		}
	}

	bestScore := -1000
	for _, move := range moveOrder {
		newBoard := b

		if !newBoard.CanPlay(move) {
			continue
		}

		newBoard.Play(move)
		branchScore := -Negamax(newBoard, -beta, -alpha)

		bestScore = max(bestScore, branchScore)
		alpha = max(alpha, branchScore)

		if alpha >= beta {
			break
		}
	}

	if bestScore <= alphaOrig {
		// upper bound
		Put(b.Key(), bestScore, UPPER)
	} else if bestScore >= beta {
		// lower bound
		Put(b.Key(), bestScore, LOWER)
	} else {
		// insert exact
		Put(b.Key(), bestScore, EXACT)
	}

	return bestScore
}
