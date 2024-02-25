package main

import (
	"fmt"

	connectfour "github.com/jacobomantilla10/connect-four"
)

func main() {
	board := connectfour.MakeBoard()
	board.DrawBoard()
	player1 := connectfour.MakePlayer(1, 'X')
	player2 := connectfour.MakePlayer(2, 'O')
	currPlayer := player2
	isWin := false
	isDraw := false
	isOver := isWin || isDraw
	for !isOver {
		if currPlayer.GetId() == 2 {
			currPlayer = player1
		} else {
			currPlayer = player2
		}
		// start a players turn and depending on whose turn it is paint different symbols on the screen
		fmt.Printf("\033[2K\rEnter column player %d: ", currPlayer.GetId())
		var col int
		_, err := fmt.Scanln(&col)

		// Insert checker into col
		for err != nil || !board.CanPlay(col-1) {
			fmt.Print("\033[1A\033[2K")
			fmt.Printf("\rInvalid insert... Enter column player %d: ", currPlayer.GetId())
			fmt.Scanln(&col)
		}

		if board.IsWinningMove(col-1, currPlayer.GetChar()) {
			isWin = true
		}

		board.Play(col-1, currPlayer.GetChar())

		isDraw = board.IsDrawn()
		isOver = isWin || isDraw
		fmt.Print("\033[15A")
		board.DrawBoard()
	}
	if isWin {
		fmt.Printf("\033[2K\rGame is over. Player %d wins!\n", currPlayer.GetId())
	} else {
		fmt.Printf("\033[2K\rGame is a draw.")
	}
}
