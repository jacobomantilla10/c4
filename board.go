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

const (
	WIDTH  = 7
	HEIGHT = 6
)

func MakeBoard() Board {
	return Board{7, 6, 0, 0, 0}
}

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

func MakeBoardFromOpening(position uint64, mask uint64, numMoves int) Board {
	return Board{7, 6, position, mask, numMoves}
}

func (b *Board) NumMoves() int {
	return b.numMoves
}

func (b *Board) Position() uint64 {
	return b.position
}

func (b *Board) Mask() uint64 {
	return b.mask
}

func (b *Board) CanPlay(y int) bool {
	return (y >= 0 && y <= 6) && b.mask&top_mask(y) == 0
}

func (b *Board) Play(y int) {
	b.position ^= b.mask
	b.mask |= (b.mask + Bottom_mask(y))
	b.numMoves++
}

func (b *Board) Key() uint64 {
	return b.position + b.mask
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
	position |= (b.mask + Bottom_mask(y)) & column_mask(y)

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
	// The game is drawn if we have played in all slots and haven't won
	return b.numMoves == HEIGHT*WIDTH
}

// Return all of the moves that a player has that don't result in a loss
func (b *Board) PossibleNonLosingMoves() []int {
	possible := b.possible()
	opponentWinningMoves := b.OpponentWinningPosition()
	// moves where opponent would win if we didn't play there
	forcedMoves := possible & opponentWinningMoves
	if forcedMoves != 0 {
		// check to see if there is more than one forced move
		if forcedMoves&(forcedMoves-1) != 0 {
			return []int{}
		} else {
			// the only possible non losing move is our forced move
			possible = forcedMoves
		}
	}
	// can't play in a slot that set's up the opponent for a connection
	res := possible & ^(opponentWinningMoves >> 1)
	// default move order used to order remaining non-losing moves
	defaultMoveOrder := [7]int{3, 2, 4, 1, 5, 0, 6}
	// our result, a slice of all of the possible non losing moves
	possibleMoves := []int{}

	// convert bitboard result into array of non-losing moves
	for _, col := range defaultMoveOrder {
		// if a column has at least one non losing move we append it to our result
		if (res>>(col*(HEIGHT+1)))&(0b0111111) != 0 {
			possibleMoves = append(possibleMoves, col)
		}
	}
	return possibleMoves
}

// Get a bitmap of all possible moves for a player
func (b *Board) possible() uint64 {
	return (b.mask + bottom_board_mask(7, 6)) & board_mask()
}

// Check all of the 3-alignments possible for a player
func (b *Board) OpponentWinningPosition() uint64 {
	position := b.position ^ b.mask

	// vertical alignment
	res := (position << 1) & (position << 2) & (position << 3)
	// horizontal alignment
	pos := (position << (HEIGHT + 1)) & (position << ((HEIGHT + 1) * 2))
	res |= pos & (position << ((HEIGHT + 1) * 3))
	res |= pos & (position >> (HEIGHT + 1))
	pos >>= ((HEIGHT + 1) * 3)
	res |= pos & (position << (HEIGHT + 1))
	res |= pos & (position >> ((HEIGHT + 1) * 3))
	// first diagonal
	pos = (position << HEIGHT) & (position << (HEIGHT * 2))
	res |= pos & (position << (HEIGHT * 3))
	res |= pos & (position >> HEIGHT)
	pos >>= (HEIGHT * 3)
	res |= pos & (position << HEIGHT)
	res |= pos & (position >> (HEIGHT * 3))
	// second diagonal
	pos = (position << (HEIGHT + 2)) & (position << ((HEIGHT + 2) * 2))
	res |= pos & (position << ((HEIGHT + 2) * 3))
	res |= pos & (position >> (HEIGHT + 2))
	pos >>= ((HEIGHT + 2) * 3)
	res |= pos & (position << (HEIGHT + 2))
	res |= pos & (position >> ((HEIGHT + 2) * 3))
	// TODO pq
	return res & (board_mask() ^ b.mask)
}

func top_mask(col int) uint64 {
	return (1 << (HEIGHT - 1)) << (col * (HEIGHT + 1))
}

func Bottom_mask(col int) uint64 {
	return 1 << (col * (HEIGHT + 1))
}

func column_mask(col int) uint64 {
	return ((1 << HEIGHT) - 1) << (col * (HEIGHT + 1))
}

func bottom_board_mask(width, height int) uint64 {
	if width == 0 {
		return 0
	}
	return bottom_board_mask(width-1, height) | uint64(1<<((width-1)*(height+1)))
}

func board_mask() uint64 {
	return 0b0111111011111101111110111111011111101111110111111
}
