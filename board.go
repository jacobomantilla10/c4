package connectfour

import "fmt"

type Board struct {
	w, h     int
	data     [6][7]rune
	numMoves int
}

func MakeBoard() Board {
	arr := [6][7]rune{}
	for i := range arr {
		arr[i] = [7]rune{' ', ' ', ' ', ' ', ' ', ' ', ' '}
	}
	return Board{6, 7, arr, 0}
}

func (b *Board) NumMoves() int {
	return b.numMoves
}

func (b *Board) CanPlay(y int) bool {
	// TODO check that the index is within the valid range to avoid panic
	return y < len(b.data[0]) && y >= 0 && b.data[0][y] == 32
}

func (b *Board) Play(y int, checker rune) {
	for x := len(b.data) - 1; x >= 0; x-- {
		if b.data[x][y] == 32 {
			b.data[x][y] = checker
			b.numMoves++
			return
		}
	}
}

func (b *Board) DrawBoard() {
	fmt.Printf("\033[2K  1   2   3   4   5   6   7  \n")
	for i := range b.data {
		fmt.Printf("\033[2K+---+---+---+---+---+---+---+\n")
		for j := range b.data[i] {
			fmt.Printf("| %c ", b.data[i][j])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("\033[2K+---+---+---+---+---+---+---+\n")
}

func (b *Board) IsWinningMove(y int, char rune) bool {
	// First figure out the row it goes into
	x := 0
	for x < len(b.data) && b.data[x][y] == 32 {
		x++
	}
	x--

	// x is now equal to our insert row
	l, r := y-1, y+1
	for l >= 0 && x >= 0 && b.data[x][l] == char {
		l--
	}
	for r < len(b.data[x]) && x >= 0 && b.data[x][r] == char {
		r++
	}
	if r-l > 4 {
		return true
	}

	h := x + 1
	for h < len(b.data) && b.data[h][y] == char {
		h++
	}
	//fmt.Printf("h: %d, x: %d\n", h, x)
	if h-x >= 4 {
		return true
	}

	// need to check up and to the left
	o, u := x-1, x+1
	l, r = y-1, y+1

	for u < len(b.data) && l >= 0 && b.data[u][l] == char {
		u++
		l--
	}
	for o >= 0 && r < len(b.data[x]) && b.data[o][r] == char {
		o--
		r++
	}
	if u-o > 4 {
		return true
	}

	o, u = x-1, x+1
	l, r = y-1, y+1

	for u < len(b.data) && r < len(b.data[x]) && b.data[u][r] == char {
		u++
		r++
	}
	for o >= 0 && l >= 0 && b.data[o][l] == char {
		o--
		l--
	}

	return u-o > 4
}

func (b *Board) IsDrawn() bool {
	for x := range b.data {
		for y := range b.data[x] {
			if b.data[x][y] == 32 {
				return false
			}
		}
	}
	return true
}
