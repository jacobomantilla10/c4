package connectfour

import (
	"fmt"
)

// TODO implement a column order array that arranges moves from the center to the outside and use that in your MiniMax
type Board struct {
	w, h     int
	position uint64
	mask     uint64
	numMoves int
}

func MakeBoard() Board {
	return Board{7, 6, 0, 0, 0}
}

// func MakeBoardWithMatrix(m [6][7]rune) Board {
// 	// TODO if this is going to be correct then we need to get the right numMoves in there and not 0
// 	return Board{6, 7, m, 0}
// }

func MakeBoardFromString(s string) (Board, error) {
	board := MakeBoard()
	//play the moves as detailed in the string
	for _, c := range s {
		col := int(c - '0')
		if !board.CanPlay(col - 1) {
			return board, fmt.Errorf("inserting into column: ")
		}
		board.Play(col - 1)
	}
	return board, nil
}

func (b *Board) NumMoves() int {
	return b.numMoves
}

func (b *Board) CanPlay(y int) bool {
	return b.position&b.top_mask(y) == 0
}

func (b *Board) top_mask(col int) uint64 {
	return 1 << (b.h - 1) << ((col - 1) * (b.h + 1))
}

func (b *Board) bottom_mask(col int) uint64 {
	return 1 << ((col - 1) * (b.h + 1))
}

func (b *Board) Play(y int) {
	b.mask |= (b.mask + b.bottom_mask(y))
	b.position ^= b.mask
	b.numMoves++
}

func (b *Board) DrawBoard() {
	pos := b.position
	mask := b.mask
	currentPlayer, opponent := 'X', 'O'
	if b.numMoves%2 == 1 {
		currentPlayer, opponent = opponent, currentPlayer
	}
	fmt.Printf("\033[2K  1   2   3   4   5   6   7  \n")
	for pos != 0 {
		fmt.Printf("\033[2K+---+---+---+---+---+---+---+\n")
		i := 0
		for i <= 7 {
			var char rune
			if mask&1 == 1 && pos&1 == 1 {
				char = currentPlayer
			} else if mask&1 == 1 && pos&1 == 0 {
				char = opponent
			} else {
				char = ' '
			}
			fmt.Printf("| %c ", char)
			fmt.Printf("|\n")
			pos = (pos >> 1)
			mask = (mask >> 1)
		}
		fmt.Printf("|\n")
		fmt.Printf("\033[2K+---+---+---+---+---+---+---+\n")
	}
}

func (b *Board) IsWinningMove(y int) bool {
	// need to add the move to the corresponding column and then do the computations on that mf
	return false
}

func (b *Board) IsDrawn() bool {
	return b.numMoves == b.h*b.w
}
