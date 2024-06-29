package solver

import (
	"github.com/jacobomantilla10/connect-four/game"
)

var DefaultMoveOrder = [7]int{3, 2, 4, 1, 5, 0, 6}
var transpositionTable = TranspositionTable{Table: make([]Transposition, 100000000), Count: 0}
var openingBook = MakeOpeningBook()

// Function used by game loop to get the best move in a position. Takes in
// a board and returns the column corresponding to the best move.
func GetBestMove(b game.Board) (int, int) {
	bestMove := 0
	bestScore := 1000

	for _, move := range DefaultMoveOrder {
		newBoard := b

		if !newBoard.CanPlay(move) {
			continue
		}

		newBoard.Play(move)
		branchScore := Negamax(newBoard, -1000, 1000, 24)
		if branchScore < bestScore {
			bestMove = move
			bestScore = branchScore
		}
	}
	return bestMove, -1 * bestScore
}

// Calls negamax function on boards using iterative deepening and null window search
// to optimize performance. Takes in a board and returns the best score for the position.
func Solve(b game.Board) int {
	// calculate the minimum and the maximum based on the current number of moves
	min := -(game.WIDTH*game.HEIGHT - b.NumMoves()) / 2
	max := (game.WIDTH*game.HEIGHT + 1 - b.NumMoves()) / 2
	// calculate the mid point between the two
	for min < max {
		mid := min + (max-min)/2
		if mid <= 0 && min/2 < mid {
			mid = min / 2
		} else if mid >= 0 && max/2 > mid {
			mid = max / 2
		}
		r := Negamax(b, mid, mid+1, 30)
		if r <= mid {
			max = r
		} else {
			min = r
		}
	}
	return min
}

// Negamax is a minimax style algorithm that plays all of the moves in a
// position (recursively until draw or end of game), computing the score for each,
// and pruning moves where it is obvious that a move couldn't benefit the current
// player. Takes in a board, alpha (the best move current player is guaranteed),
// beta (the best move opposite player is guaranted). Returns the score of the board.
// The score is the amount of moves from the last move that a player will win in: 1
// means that the current player will win in the last move, -5 means they'll lose in
// the 5th to last move, and 0 is a draw.
func Negamax(b game.Board, alpha, beta, depth int) int {
	alphaOrig := alpha

	if b.IsDrawn() {
		return 0
	}

	// loop through move order to see if an immediate win is available
	for _, move := range DefaultMoveOrder {
		if b.CanPlay(move) && b.IsWinningMove(move) {
			return ((game.WIDTH*game.HEIGHT + 1 - b.NumMoves()) / 2)
		}
	}

	if b.NumMoves() == 12 {
		opening := openingBook.Get_book_value(b)
		if opening != -300 {
			return int(opening)
		}
	}

	// Check to see if opponent has a move that will cause us to lose next turn
	possibleNonLosingMoves := b.PossibleNonLosingMoves()
	if len(possibleNonLosingMoves) == 0 {
		return -(game.WIDTH*game.HEIGHT - b.NumMoves()) / 2
	}

	// Since we know we can't win with this move, we know our best score is bounded
	best := ((game.WIDTH*game.HEIGHT - 1 - b.NumMoves()) / 2)
	// If our best is worse than beta, our opponents best guaranteed move (beta) gets set to best
	beta = min(beta, best)
	if beta <= alpha {
		return beta
	}

	// Since we know our opponent can't win with their next move, we know their best score,
	// and therefore our worst score, is bounded.
	worst := -(game.WIDTH*game.HEIGHT - 2 - b.NumMoves()) / 2
	// If our worst score is better than alpha, our best guaranteed move (alpha) gets set to worst
	alpha = max(alpha, worst)
	if beta <= alpha {
		return alpha
	}

	// Check to see if we have stored current position in transposition table
	tt := transpositionTable.Get(b.Key())
	if tt.Val != -128 {
		if tt.Flag == EXACT {
			return int(tt.Val)
		} else if tt.Flag == LOWER {
			alpha = max(alpha, int(tt.Val))
		} else if tt.Flag == UPPER {
			beta = min(beta, int(tt.Val))
		}

		if beta <= alpha {
			return int(tt.Val)
		}
	}

	// Create ordered moves struct and fill it with correctly ordered moves
	moveOrder := OrderedMoves()
	for _, move := range possibleNonLosingMoves {
		newBoard := b
		newBoard.Play(move)
		score := PopCount(newBoard.OpponentWinningPosition())
		moveOrder.Insert(move, score)
	}

	// if we get to a depth of 0 we return the score (population count) of the best move
	if depth == 0 {
		return moveOrder.moves[0].score
	}

	// Compute score for position and save it into bestScore
	bestScore := worst
	for _, move := range moveOrder.moves {
		newBoard := b

		if !newBoard.CanPlay(move.col) {
			continue
		}
		// Play move on new board
		newBoard.Play(move.col)
		// Compute negamax score for our opponent on new board
		branchScore := -Negamax(newBoard, -beta, -alpha, depth-1)
		// Update best score and alpha if new found score is greater
		bestScore = max(bestScore, branchScore)
		alpha = max(alpha, branchScore)

		if alpha >= beta {
			break
		}
	}

	// Store computed score into transposition table
	if bestScore <= alphaOrig {
		// We never pushed alpha up, so every move we searched pruned.
		// we don't know the true value, but we know it's an upper bound.
		transpositionTable.Put(b.Key(), bestScore, UPPER)
	} else if bestScore >= beta {
		// Alpha was pushed up, but we pruned so all we know is that
		// the score we have is the best of all the moves we examined.
		// therefore we can store is as a lower bound
		transpositionTable.Put(b.Key(), bestScore, LOWER)
	} else {
		// We pushed alpha up, but we didn't prune, so we went through
		// every move and found a new best move. We know that this
		// score is the exact score and so we store is as exact.
		transpositionTable.Put(b.Key(), bestScore, EXACT)
	}

	return bestScore
}
