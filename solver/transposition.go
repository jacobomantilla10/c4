package solver

// Transposition object cointains a key that is an encoding of the
// position, a value that is the score of that position, and a
// flag that denotes whether the score is an upper bound, a lower
// bound, or an exact score.
type Transposition struct {
	Key  uint64
	Flag uint8
	Val  int8
}

const (
	UPPER uint8 = iota
	LOWER uint8 = iota
	EXACT uint8 = iota
)

// TranspositionTable is made up of a slice of Transpositions and
// a count that tells us how many positions we have stored.
type TranspositionTable struct {
	Table []Transposition
	Count int
}

// Gets the index of a key in the transposition table
func index(key uint64, size int) int {
	return int(key) % size
}

// Inserts a position into the transposition table
func (t *TranspositionTable) Put(key uint64, val int, flag uint8) {
	// calculate the index based on the position
	index := index(key, len(t.Table))
	t.Table[index].Key = key
	t.Table[index].Val = int8(val)
	t.Table[index].Flag = flag
	t.Count++
}

// Gets and returns a position given by the key from the transposition
// table.
func (t *TranspositionTable) Get(key uint64) Transposition {
	index := index(key, len(t.Table))
	if t.Table[index].Key == key {
		return t.Table[index]
	} else {
		return Transposition{Key: 0, Val: int8(-128), Flag: 0}
	}
}
