package main

import "fmt"

func main() {
	board := makeBoard()
	board.DrawBoard()
	turn := 2

	// Add turns by adding players
	isOver := false
	for !isOver {
		if turn == 2 {
			turn = 1
		} else {
			turn = 2
		}
		// start a players turn and depending on whose turn it is paint different symbols on the screen
		fmt.Printf("\033[2K\rEnter column player %d: ", turn)
		var col int
		fmt.Scanln(&col)
		// Insert checker into col
		err := board.InsertIntoCol(col-1, turn)
		for err != nil {
			fmt.Print("\033[1A\033[2K")
			fmt.Printf("\rInvalid insert... Enter column player %d: ", turn)
			fmt.Scanln(&col)
			err = board.InsertIntoCol(col-1, turn)
		}
		isWin := board.IsWin()
		isDraw := board.IsDrawn()
		isOver = isWin || isDraw
		fmt.Print("\033[15A")
		board.DrawBoard()
	}
	fmt.Printf("\033[2K\rGame is over. Player %d wins!\n", turn)
}

type Board struct {
	w, h int
	data [6][7]int
}

func makeBoard() Board {
	arr := [6][7]int{}
	return Board{6, 7, arr}
}

func (b *Board) At(x, y int) int {
	return b.data[x][y]
}

func (b *Board) Set(x, y, new int) {
	b.data[x][y] = new
}

func (b *Board) InsertIntoCol(y, checker int) error {
	if y > len(b.data) || y < 0 {
		return fmt.Errorf("invalid insert")
	}

	for x := len(b.data) - 1; x >= 0; x-- {
		if b.data[x][y] == 0 {
			b.Set(x, y, checker)
			return nil
		}
	}
	return fmt.Errorf("invalid insert")
}

func (b *Board) DrawBoard() {
	fmt.Printf("\033[2K  1   2   3   4   5   6   7  \n")
	for i := range b.data {
		fmt.Printf("\033[2K+---+---+---+---+---+---+---+\n")
		for j := range b.data[i] {
			fmt.Printf("| %d ", b.data[i][j])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("\033[2K+---+---+---+---+---+---+---+\n")
}

func (b *Board) IsWin() bool {
	for x := len(b.data) - 1; x >= 0; x-- {
		// C is the index of the first occurrence of a non-empty checker in a horizontal
		c := 0
		for y := range b.data[x] {
			// Check horizontal
			if b.data[x][y] != b.data[x][c] || b.data[x][y] == 0 {
				c = y
			}
			if y-c >= 3 {
				return true
			}
			// Check vertical
			h := x - 1
			for h >= 0 && b.data[h][y] == b.data[x][y] && b.data[h][y] != 0 {
				h--
				if x-h == 4 {
					return true
				}
			}
			// Check left up diagonal
			h = x - 1
			w := y - 1
			for h >= 0 && w >= 0 && b.data[h][w] == b.data[x][y] && b.data[h][w] != 0 {
				h--
				w--
				if x-h == 4 {
					return true
				}
			}
			// Check up right diagonal
			h = x - 1
			w = y + 1
			for h >= 0 && w < len(b.data[x]) && b.data[h][w] == b.data[x][y] && b.data[h][w] != 0 {
				h--
				w++
				if x-h == 4 {
					return true
				}
			}
		}
	}
	return false
}

func (b *Board) IsDrawn() bool {
	for x := range b.data {
		for y := range b.data[x] {
			if b.data[x][y] == 0 {
				return false
			}
		}
	}
	return true
}
