package main

import "fmt"

func main() {
	board := makeBoard()
	board.DrawBoard()
	player1 := makePlayer(1, 'X')
	player2 := makePlayer(2, 'O')
	currPlayer := player2
	isWin := false
	isDraw := false
	isOver := isWin || isDraw
	for !isOver {
		if currPlayer.playerId == 2 {
			currPlayer = player1
		} else {
			currPlayer = player2
		}
		// start a players turn and depending on whose turn it is paint different symbols on the screen
		fmt.Printf("\033[2K\rEnter column player %d: ", currPlayer.playerId)
		var col int
		fmt.Scanln(&col)
		// Insert checker into col
		err := board.InsertIntoCol(col-1, currPlayer.character)
		for err != nil {
			fmt.Print("\033[1A\033[2K")
			fmt.Printf("\rInvalid insert... Enter column player %d: ", currPlayer.playerId)
			fmt.Scanln(&col)
			err = board.InsertIntoCol(col-1, currPlayer.character)
		}
		isWin = board.IsWin()
		isDraw = board.IsDrawn()
		isOver = isWin || isDraw
		fmt.Print("\033[15A")
		board.DrawBoard()
	}
	if isWin {
		fmt.Printf("\033[2K\rGame is over. Player %d wins!\n", currPlayer.playerId)
	} else {
		fmt.Printf("\033[2K\rGame is a draw.")
	}
}

type Player struct {
	playerId  int
	character rune
}

func makePlayer(id int, symbol rune) Player {
	return Player{id, symbol}
}

type Board struct {
	w, h int
	data [6][7]rune
}

func makeBoard() Board {
	arr := [6][7]rune{}
	for i := range arr {
		arr[i] = [7]rune{' ', ' ', ' ', ' ', ' ', ' ', ' '}
	}
	return Board{6, 7, arr}
}

func (b *Board) At(x, y int) rune {
	return b.data[x][y]
}

func (b *Board) Set(x, y int, new rune) {
	b.data[x][y] = new
}

func (b *Board) InsertIntoCol(y int, checker rune) error {
	if y > len(b.data) || y < 0 {
		return fmt.Errorf("invalid insert")
	}

	for x := len(b.data) - 1; x >= 0; x-- {
		if b.data[x][y] == 32 {
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
			fmt.Printf("| %c ", b.data[i][j])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("\033[2K+---+---+---+---+---+---+---+\n")
}

func (b *Board) IsWin() bool {
	// Figure out how to check for empty runes
	// What is the nil value for a rune?
	for x := len(b.data) - 1; x >= 0; x-- {
		// C is the index of the first occurrence of a non-empty checker in a horizontal
		c := 0
		for y := range b.data[x] {
			// Check horizontal
			if b.data[x][y] != b.data[x][c] || b.data[x][y] == 32 {
				c = y
			}
			if y-c >= 3 {
				return true
			}
			// Check vertical
			h := x - 1
			for h >= 0 && b.data[h][y] == b.data[x][y] && b.data[h][y] != 32 {
				h--
				if x-h == 4 {
					return true
				}
			}
			// Check left up diagonal
			h = x - 1
			w := y - 1
			for h >= 0 && w >= 0 && b.data[h][w] == b.data[x][y] && b.data[h][w] != 32 {
				h--
				w--
				if x-h == 4 {
					return true
				}
			}
			// Check up right diagonal
			h = x - 1
			w = y + 1
			for h >= 0 && w < len(b.data[x]) && b.data[h][w] == b.data[x][y] && b.data[h][w] != 32 {
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
			if b.data[x][y] == 32 {
				return false
			}
		}
	}
	return true
}
