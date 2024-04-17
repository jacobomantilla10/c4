package solver

type move struct {
	col   int
	score int
}

type orderedMoves struct {
	moves [7]move
	size  int
}

func OrderedMoves() orderedMoves {
	return orderedMoves{[7]move{}, 0}
}

func (o *orderedMoves) Insert(col, score int) {
	o.size++
	i := o.size - 1
	for ; i >= 1 && o.moves[i-1].score < score; i-- {
		o.moves[i] = o.moves[i-1]
	}
	o.moves[i] = move{col: col, score: score}
}

func PopCount(bitboard uint64) int {
	count := 0
	for bitboard != 0 {
		count++
		bitboard &= bitboard - 1
	}
	// get the count of 1s in the bitboard
	return count
}
