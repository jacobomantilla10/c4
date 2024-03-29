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

var TranspositionTable = make([]Transposition, 10000)

// implement index put and get
func index(key uint64) uint64 {
	return uint64(key % 10000)
}

func Put(key uint64, val, flag int) {
	// calculate the index based on the position
	index := index(key)
	TranspositionTable[index].Key = key
	TranspositionTable[index].Val = val
	TranspositionTable[index].Flag = flag
}

func Get(key uint64) Transposition {
	index := index(key)
	if TranspositionTable[index].Key == key {
		return TranspositionTable[index]
	} else {
		return Transposition{Key: 0, Val: -999, Flag: 0}
	}
}
