package solver

import (
	connectfour "github.com/jacobomantilla10/connect-four"
)

var defaultMoveOrder = [7]int{3, 2, 4, 1, 5, 0, 6}
var transpositionTable = TranspositionTable{Table: make([]Transposition, 1000), Count: 0}

func GetBestMove(b connectfour.Board) int {
	bestMove := 0
	bestScore := 1000

	for _, move := range defaultMoveOrder {
		if b.CanPlay(move) && b.IsWinningMove(move) {
			return move
		}
	}

	for _, move := range defaultMoveOrder {
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

func solve(b connectfour.Board) int {
	// calculate the minimum and the maximum based on the current number of moves
	min := -(connectfour.WIDTH*connectfour.HEIGHT - b.NumMoves()) / 2
	max := (connectfour.WIDTH*connectfour.HEIGHT + 1 - b.NumMoves()) / 2
	// calculate the mid point between the two
	for min < max {
		mid := min + (max-min)/2
		if mid <= 0 && min/2 < mid {
			mid = min / 2
		} else if mid >= 0 && max/2 > mid {
			mid = max / 2
		}
		r := Negamax(b, mid, mid+1)
		if r <= mid {
			max = r
		} else {
			min = r
		}
	}
	return min
}

func Negamax(b connectfour.Board, alpha, beta int) int {
	alphaOrig := alpha

	// TODO query transpositiontable to get value and make check the flag

	if b.IsDrawn() {
		return 0
	}

	for _, move := range defaultMoveOrder {
		if b.CanPlay(move) && b.IsWinningMove(move) {
			return ((connectfour.WIDTH*connectfour.HEIGHT + 1 - b.NumMoves()) / 2)
		}
	}

	// If there are no moves that don't cause us the opponent to win return opponent winning score
	if possibleNonLosingMoves := b.PossibleNonLosingMoves(); possibleNonLosingMoves == 0 {
		return -(connectfour.WIDTH*connectfour.HEIGHT - b.NumMoves()) / 2
	}

	// TODO update bestscore to be computed off of b.NumMoves()
	best := ((connectfour.WIDTH*connectfour.HEIGHT - 1 - b.NumMoves()) / 2)
	beta = min(beta, best)
	if beta <= alpha {
		return beta
	}

	worst := -(connectfour.WIDTH*connectfour.HEIGHT - 2 - b.NumMoves()) / 2
	alpha = max(alpha, worst)
	if beta <= alpha {
		return alpha
	}

	tt := transpositionTable.Get(b.Key())
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

	// TODO write code that uses the sorter code to generate better move ordering
	moveOrder := OrderedMoves()
	for _, move := range defaultMoveOrder {
		// get bitboard and pass it into population count to get the score
		// the pass in the score and the column into the insert function
		newBoard := b
		newBoard.Play(move)
		score := PopCount(newBoard.OpponentWinningPosition())
		moveOrder.Insert(move, score)
	}

	bestScore := -1000
	for _, move := range moveOrder.moves {
		newBoard := b

		if !newBoard.CanPlay(move.col) {
			continue
		}

		newBoard.Play(move.col)
		branchScore := -Negamax(newBoard, -beta, -alpha)

		bestScore = max(bestScore, branchScore)
		alpha = max(alpha, branchScore)

		if alpha >= beta {
			break
		}
	}

	if bestScore <= alphaOrig {
		// upper bound
		transpositionTable.Put(b.Key(), bestScore, UPPER)
	} else if bestScore >= beta {
		// lower bound
		transpositionTable.Put(b.Key(), bestScore, LOWER)
	} else {
		// insert exact
		transpositionTable.Put(b.Key(), bestScore, EXACT)
	}

	return bestScore
}
