package solver

import (
	connectfour "github.com/jacobomantilla10/connect-four"
)

// TODO add function to emulate move from AI calling MiniMax
func GetBestMove(b connectfour.Board) int {
	bestMove := 0
	bestScore := -1000

	for i := 0; i <= 6; i++ {
		if b.CanPlay(i) && b.IsWinningMove(i) {
			return i
		}
	}

	for i := 0; i <= 6; i++ {
		newBoard := b

		if !newBoard.CanPlay(i) {
			continue
		}

		newBoard.Play(i)
		branchScore := MiniMax(newBoard, 14, -1000, 1000, 1)
		if branchScore > bestScore {
			bestMove = i
			bestScore = branchScore
		}
	}
	return bestMove
}

func MiniMax(b connectfour.Board, depth, alpha, beta, isMaximizingPlayer int) int {
	// check to see if there is a draw and if there is return 0 and not sure what move //TODO ??
	if b.IsDrawn() || depth == 0 {
		return 0
	}

	// check to see if there is any immediate win and if there is return the computed value from it as well as the column associated with it
	for i := 0; i <= 6; i++ {
		if b.CanPlay(i) && b.IsWinningMove(i) {
			return isMaximizingPlayer * ((7*6 + 1 - b.NumMoves()) / 2)
		}
	}

	var bestScore int
	if isMaximizingPlayer == 1 {
		bestScore = -1000
		for i := 0; i <= 6; i++ {
			newBoard := b

			if !newBoard.CanPlay(i) {
				continue
			}

			newBoard.Play(i)
			branchScore := MiniMax(newBoard, depth-1, alpha, beta, -1)

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
		for i := 0; i <= 6; i++ {
			newBoard := b

			if !newBoard.CanPlay(i) {
				continue
			}

			newBoard.Play(i)
			branchScore := MiniMax(newBoard, depth-1, alpha, beta, 1)

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
