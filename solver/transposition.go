package solver

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

type TranspositionTable struct {
	Table []Transposition
	Count int
}

// implement index put and get
func index(key uint64, size int) int {
	return int(key) % size
}

func (t *TranspositionTable) Put(key uint64, val int, flag uint8) {
	// calculate the index based on the position
	index := index(key, len(t.Table))
	t.Table[index].Key = key
	t.Table[index].Val = int8(val)
	t.Table[index].Flag = flag
}

func (t *TranspositionTable) Get(key uint64) Transposition {
	index := index(key, len(t.Table))
	if t.Table[index].Key == key {
		return t.Table[index]
	} else {
		return Transposition{Key: 0, Val: int8(-128), Flag: 0}
	}
}
