package main

import "fmt"

func emptyBoard() [7][6]string {
	return [7][6]string{
		{"Empty", "Empty", "Empty", "Empty", "Empty", "Empty"},
		{"Empty", "Empty", "Empty", "Empty", "Empty", "Empty"},
		{"Empty", "Empty", "Empty", "Empty", "Empty", "Empty"},
		{"Empty", "Empty", "Empty", "Empty", "Empty", "Empty"},
		{"Empty", "Empty", "Empty", "Empty", "Empty", "Empty"},
		{"Empty", "Empty", "Empty", "Empty", "Empty", "Empty"},
		{"Empty", "Empty", "Empty", "Empty", "Empty", "Empty"},
	}
}

func boardFromString(arr [7][6]string, board string) [7][6]string {
	for i, ch := range board {
		// convert the current utf character into a number and subtract one to find the column we need to insert into
		// loop into first open spot in array and add the correct value depending on parity of turn
		insertCol := int(ch-'0') - 1
		insertRow := 0
		insertVal := "Player1"
		if i%2 == 1 {
			insertVal = "Player2"
		}
		for j := 0; arr[insertCol][j] != "Empty" && j < 6; j++ {
			insertRow++
		}
		if insertRow != 6 {
			arr[insertCol][insertRow] = insertVal
		}
	}
	return arr
}

func formatOutcome(score, numMoves int) string {
	var winner string
	var victoryMove int
	if score < 0 {
		winner = "Red"
		victoryMove = 22 + score
	} else {
		winner = "Yellow"
		victoryMove = 22 - score
	}
	movesByPlayer := (numMoves + 1) / 2
	numMovesToVictory := victoryMove - movesByPlayer

	var outcome string
	if numMovesToVictory == 1 {
		outcome = fmt.Sprintf("%s wins in %d move", winner, numMovesToVictory)
	} else {
		outcome = fmt.Sprintf("%s wins in %d moves", winner, numMovesToVictory)
	}
	return outcome
}
