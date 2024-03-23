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
	return b.mask&b.top_mask(y) == 0
}

func (b *Board) top_mask(col int) uint64 {
	return (1 << (b.h - 1)) << (col * (b.h + 1))
}

func (b *Board) bottom_mask(col int) uint64 {
	return 1 << (col * (b.h + 1))
}

func (b *Board) column_mask(col int) uint64 {
	return ((1 << b.h) - 1) << (col * (b.h + 1))
}

func (b *Board) Play(y int) {
	b.position ^= b.mask
	b.mask |= (b.mask + b.bottom_mask(y))
	// b.position ^= b.mask
	b.numMoves++
}

func (b *Board) DrawBoard() {
	// write algorithm to convert number to rune array used to render position
	//posArr := []uint64{0, 0, 0, 0, 0, 0, 0} // need to map bit at (i, j) to (j, i)
	currentPlayer, opponent := 'O', 'X'
	if b.numMoves%2 == 1 {
		currentPlayer, opponent = opponent, currentPlayer
	}
	posArr := [6][7]rune{
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{' ', ' ', ' ', ' ', ' ', ' ', ' '},
	}
	k := 0
	for i := 5; i >= 0; i-- {
		for j := 0; j < len(posArr[i]); j++ {
			currPos := b.position >> k
			currMask := b.mask >> k
			var char rune
			if (currMask>>(j*6))&1 == 1 && (currPos>>(j*6))&1 == 1 {
				char = currentPlayer
			} else if (currMask>>(j*6))&1 == 1 && (currPos>>(j*6))&1 == 0 {
				char = opponent
			} else {
				char = ' '
			}
			posArr[i][j] = char
		}
		k++
	}
	fmt.Printf("\033[2K  1   2   3   4   5   6   7  \n")
	fmt.Printf("\033[2K+---+---+---+---+---+---+---+\n")
	for i := range posArr {
		for j := range posArr[i] {
			fmt.Printf("| %c ", posArr[i][j])
		}
		fmt.Printf("|\n")
		fmt.Printf("\033[2K+---+---+---+---+---+---+---+\n")
		i++
	}
}

func (b *Board) IsWinningMove(y int) bool {
	// need to add the move to the corresponding column and then do the computations on that mf	return false
	position := b.position
	position |= (b.mask + b.bottom_mask(y)) & b.column_mask(y)

	// now that you have all of the information you need to calculate all of the alignments

	// horizontal
	pos := (position << (b.h + 1)) & position
	if (pos<<((b.h+1)*2))&pos != 0 {
		return true
	}

	pos = (position << (b.h)) & position
	if (pos<<(b.h*2))&pos != 0 {
		return true
	}

	pos = (position << (b.h + 2)) & position
	if (pos<<((b.h+2)*2))&pos != 0 {
		return true
	}

	pos = (position << 1) & position
	if (pos<<2)&pos != 0 {
		return true
	}

	return false
}

func (b *Board) IsDrawn() bool {
	return b.numMoves == b.h*b.w
}
