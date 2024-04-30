package solver

// Represents a move, where col is the column corresponding to a move
// and score is the amount of three alignments the move creates. The higher
// the score the better.
type move struct {
	col   int
	score int
}

// orderedMoves contains the list of the 7 possible moves in a position,
// ordered descreasingly by how many three alignments they threaten to create.
// It also contains size, which contains the amount of items in the array at a given time.
type orderedMoves struct {
	moves [7]move
	size  int
}

// Creates and returns empty orderedMoves struct
func OrderedMoves() orderedMoves {
	return orderedMoves{[7]move{}, 0}
}

// Inserts move into orderedMoves array in sorted (by amount of threats it generates) order
func (o *orderedMoves) Insert(col, score int) {
	o.size++
	i := o.size - 1
	for ; i >= 1 && o.moves[i-1].score < score; i-- {
		o.moves[i] = o.moves[i-1]
	}
	o.moves[i] = move{col: col, score: score}
}

// Popcount computes and returns the score of the move (the amount of three-alignments it generates)
func PopCount(bitboard uint64) int {
	count := 0
	for bitboard != 0 {
		count++
		bitboard &= bitboard - 1
	}
	return count
}
