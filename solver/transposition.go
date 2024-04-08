package solver

type Transposition struct {
	Key  uint64
	Flag int
	Val  int
}

const (
	UPPER = iota
	LOWER = iota
	EXACT = iota
)

type TranspositionTable struct {
	Table []Transposition
	Count int
}

// implement index put and get
func index(key uint64, size int) int {
	return int(key) % size
}

func (t *TranspositionTable) Put(key uint64, val, flag int) {
	// calculate the index based on the position
	index := index(key, len(t.Table))
	t.Table[index].Key = key
	t.Table[index].Val = val
	t.Table[index].Flag = flag
}

func (t *TranspositionTable) Get(key uint64) Transposition {
	index := index(key, len(t.Table))
	if t.Table[index].Key == key {
		return t.Table[index]
	} else {
		return Transposition{Key: 0, Val: -999, Flag: 0}
	}
}
