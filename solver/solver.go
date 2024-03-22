package solver

import (
	connectfour "github.com/jacobomantilla10/connect-four"
)

var moveOrder = [7]int{3, 2, 4, 1, 5, 0, 6} // Optimal move exploration order. We loop over this to explore optimal branches first.

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
		branchScore := MiniMax(newBoard, -1000, 1000, 1)
		if branchScore < bestScore {
			bestMove = move
			bestScore = branchScore
		}
	}
	return bestMove
}

func MiniMax(b connectfour.Board, alpha, beta, isMaximizingPlayer int) int {
	// check to see if there is a draw and if there is return 0 and not sure what move //TODO ??
	if b.IsDrawn() {
		return 0
	}

	// check to see if there is any immediate win and if there is return the computed value from it as well as the column associated with it
	for _, move := range moveOrder {
		if b.CanPlay(move) && b.IsWinningMove(move) {
			return isMaximizingPlayer * ((7*6 + 1 - b.NumMoves()) / 2)
		}
	}

	var bestScore int
	if isMaximizingPlayer == 1 {
		bestScore = -1000
		for _, move := range moveOrder {
			newBoard := b

			if !newBoard.CanPlay(move) {
				continue
			}

			newBoard.Play(move)
			branchScore := MiniMax(newBoard, alpha, beta, -1)

			if branchScore > bestScore {
				bestScore = branchScore
			}

			if branchScore > alpha {
				alpha = branchScore
			}

			if beta <= alpha {
				break
			}
		}
	} else {
		bestScore = 1000
		for _, move := range moveOrder {
			newBoard := b

			if !newBoard.CanPlay(move) {
				continue
			}

			newBoard.Play(move)
			branchScore := MiniMax(newBoard, alpha, beta, 1)

			if branchScore < bestScore {
				bestScore = branchScore
			}

			if branchScore < beta {
				beta = branchScore
			}

			if beta <= alpha {
				break
			}
		}
	}
	return bestScore
}
