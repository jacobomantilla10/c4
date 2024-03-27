package solver

import (
	connectfour "github.com/jacobomantilla10/connect-four"
)

var moveOrder = [7]int{3, 2, 4, 1, 5, 0, 6} // Optimal move exploration order. We loop over this to explore optimal branches first.
// TODO create transposition table type that is based on hashmap from uint64 to uint8 that stores the value of the position
// TODO then we use the transposition table to check to see if we've cached the result and don't have to compute the minimax
// TODO we need then at the end to add the value that was computed as optimal for a position to the transposition table
// TODO we need to figure out a way to keep the transposition table to a certain size when inserting and figure out what to remove
// note on the point above: we'd probably want to start removing cached results from least added to maximize performance
type TranspositionTable struct {
	table map[uint64]int
	// indices []uint64
}

func (t *TranspositionTable) get(i uint64) int {
	val, ok := t.table[i]
	if !ok {
		return -999
	}
	return val
}

func (t *TranspositionTable) set(i uint64, val int) {
	t.table[i] = val
}

// var transposition = TranspositionTable{
// 	table: map[uint64]int{},
// }

func GetBestMove(b connectfour.Board) int {
	bestMove := 0
	bestScore := 1000
	var transposition = TranspositionTable{
		table: map[uint64]int{},
	}

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
		branchScore := MiniMax(newBoard, -1000, 1000, 1, transposition)
		if branchScore < bestScore {
			bestMove = move
			bestScore = branchScore
		}
	}
	return bestMove
}

func MiniMax(b connectfour.Board, alpha, beta, isMaximizingPlayer int, t TranspositionTable) int {
	if b.IsDrawn() {
		return 0
	}

	// check to see if there is any immediate win and if there is return the computed value from it
	for _, move := range moveOrder {
		if b.CanPlay(move) && b.IsWinningMove(move) {
			return isMaximizingPlayer * ((7*6 + 1 - b.NumMoves()) / 2)
		}
	}

	var bestScore int
	if isMaximizingPlayer == 1 {
		bestScore = -1000

		cachedScore := t.get(b.Key())
		if cachedScore != -999 {
			if cachedScore > alpha {
				alpha = cachedScore
			}

			if beta <= alpha {
				return cachedScore
			}
		}
		for _, move := range moveOrder {
			newBoard := b

			if !newBoard.CanPlay(move) {
				continue
			}

			newBoard.Play(move)

			branchScore := MiniMax(newBoard, alpha, beta, -1, t)

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

		cachedScore := t.get(b.Key())
		if cachedScore != -999 {
			if cachedScore < beta {
				beta = cachedScore
			}

			if beta <= alpha {
				return cachedScore
			}
		}
		for _, move := range moveOrder {
			newBoard := b

			if !newBoard.CanPlay(move) {
				continue
			}

			newBoard.Play(move)

			branchScore := MiniMax(newBoard, alpha, beta, 1, t)

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
	t.set(b.Key(), bestScore)
	return bestScore
}

// 0000000  0000000
// 1010000  1111100 XOXOO..
// 1001100  1111111 XOOXXOO
// 0011010  1111111 OOXXOXO
// 1100101  1111111 XXOOXOX
// 0111001  1111111 OXXXOOX
// 1001110  1111111 XOOXXXO
// this position gives us both 0 and -1
// the way this happens is I get the position and it's not in the map, so I pass it to minimizing player, they can win in 1 so they return -1 and we set it to that
// then the next time around I look in the map and I get -1 as the best move in that position but it's actually 0 that is the best move not -1

// 11111

// 0000000 0000000 .......
// 0000000 0000000 .......
// 0000001 1001001 O..O..X
// 1100101 1101101 XX.OX.X
// 0011010 1111111 OOXXOXO
// 1000110 1111111 XOOOXXO
// 0110001 1111111 OXXOOOX
