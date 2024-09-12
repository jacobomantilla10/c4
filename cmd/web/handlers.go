package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jacobomantilla10/c4/internal/game"
	"github.com/jacobomantilla10/c4/internal/solver"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	tpl := app.templateCache["home"]

	data := newTemplateData()

	err := tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func (app *application) restart(w http.ResponseWriter, r *http.Request) {
	tpl := app.templateCache["board"]

	data := newTemplateData()

	err := tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func (app *application) getNextMove(w http.ResponseWriter, r *http.Request) {
	board := r.URL.Query().Get("Board")
	move := r.URL.Query().Get("Move")

	board += move

	tpl := app.templateCache["board"]

	fmt.Println("Board: ", board)
	fmt.Println("Move: ", move)

	// use engine to calculate best move given board string and save the result
	// in new variable which you append to board
	b, err := game.MakeBoardFromString(board)
	if err != nil {
		panic(err)
	}
	insertCol, _ := solver.GetBestMove(b)
	board += strconv.Itoa(insertCol + 1)
	boardArr := boardFromString(emptyBoard(), board)

	data := newTemplateData()
	data.Board = boardArr
	data.BoardString = board
	data.IsGameOver = b.IsWinningMove(insertCol)

	err = tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
