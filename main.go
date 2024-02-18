package main

import "fmt"

func main() {
	board := makeBoard(7, 6)
	board.DrawBoard()
	// isOver := false
	// for !isOver {

	// }
	// Loop until game is over
	// Take input from user
	// Drop token in correct place using Set
	// Draw board again using drawboard and updated board
	// Check to see if that particular token won the game and update isOver var
	// If game isn't over loop again
	// If game is over break out of loop and let user know that someone already won
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

func (b *Board) DrawBoard() {
	fmt.Printf("  1   2   3   4   5   6   7  \n")
	for i := range b.data {
		fmt.Printf("+---+---+---+---+---+---+---+\n")
		for j := range b.data[i] {
			fmt.Printf("| %d ", b.data[i][j])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("+---+---+---+---+---+---+---+\n")
}
