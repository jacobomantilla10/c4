package main

import "fmt"

func main() {
	board := makeBoard(7, 6)
	board.DrawBoard()

	isOver := false
	for !isOver {
		fmt.Printf("\033[2K\rEnter column: ")
		var col int
		fmt.Scanln(&col)
		err := board.InsertIntoCol(col - 1)
		for err != nil {
			fmt.Print("\033[1A\033[2K")
			fmt.Printf("\rInvalid insert... Enter column: ")
			fmt.Scanln(&col)
			err = board.InsertIntoCol(col - 1)
		}
		// TODO Don't need to go all the way up to home... Just the next line below where the program was run from
		fmt.Print("\033[15A")
		board.DrawBoard()
	}
	// Loop until game is over

	// Drop token in correct place using Set
	// Draw board again using drawboard and updated board
	// Check to see if that particular token won the game and update isOver var
	// If game isn't over loop again
	// If game is over break out of loop and let user know that someone already won
	// Consider also that game can end in a draw if there are a no more legal moves
}

type Board struct {
	w, h int
	data [6][7]int
}

func makeBoard(w, h int) Board {
	arr := [6][7]int{}
	return Board{w, h, arr}
}

func (b *Board) At(x, y int) int {
	return b.data[x][y]
}

func (b *Board) Set(x, y, new int) {
	b.data[x][y] = new
}

func (b *Board) InsertIntoCol(y int) error {
	// need to check to see if there already is a symbol in that spot
	// if there is then you need to go to the next spot
	// if there isn't then you can insert it into that spot
	// need to loop through the rows backwards and only check the c position
	if y > len(b.data) || y < 0 {
		return fmt.Errorf("invalid insert")
	}

	for x := len(b.data) - 1; x >= 0; x-- {
		if b.data[x][y] == 0 {
			b.Set(x, y, 1)
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
