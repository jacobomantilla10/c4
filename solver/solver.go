package solver

import (
	connectfour "github.com/jacobomantilla10/connect-four"
)

// TODO add function to emulate move from AI calling MiniMax
func GetBestMove(b connectfour.Board) int {
	bestMove := 0
	bestScore := -1000

	for i := 0; i <= 6; i++ {
		if b.CanPlay(i) && b.IsWinningMove(i, 'O') {
			return i
		}
	}

	for i := 0; i <= 6; i++ {
		newBoard := b

		if !newBoard.CanPlay(i) {
			continue
		}

		newBoard.Play(i, 'O')
		branchScore := MiniMax(newBoard, 14, -1000, 1000, 'X')
		if branchScore > bestScore {
			bestMove = i
			bestScore = branchScore
		}
	}
	return bestMove
}

func MiniMax(b connectfour.Board, depth, alpha, beta int, checker rune) int {
	multipliers := map[rune]int{'X': -1, 'O': 1}
	// check to see if there is a draw and if there is return 0 and not sure what move //TODO ??
	if b.IsDrawn() || depth == 0 {
		return 0
	}

	// check to see if there is any immediate win and if there is return the computed value from it as well as the column associated with it
	for i := 0; i <= 6; i++ {
		if b.CanPlay(i) && b.IsWinningMove(i, checker) {
			return multipliers[checker] * ((7*6 + 1 - b.NumMoves()) / 2)
		}
	}

	var bestScore int
	if checker == 'X' {
		bestScore = 1000
		for i := 0; i <= 6; i++ {
			newBoard := b

			if !newBoard.CanPlay(i) {
				continue
			}

			newBoard.Play(i, checker)
			branchScore := MiniMax(newBoard, depth-1, alpha, beta, 'O')

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
	} else {
		bestScore = -1000
		for i := 0; i <= 6; i++ {
			newBoard := b
			newBoard.Play(i, checker)

			if !newBoard.CanPlay(i) {
				continue
			}

			branchScore := MiniMax(newBoard, depth-1, alpha, beta, 'X')

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
	}
	return bestScore
}
