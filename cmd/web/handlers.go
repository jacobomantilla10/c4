package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/jacobomantilla10/c4/internal/game"
	"github.com/jacobomantilla10/c4/internal/solver"
)

func home(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("../../ui/html/home.gohtml")
	if err != nil {
		panic(err)
	}
	boardArr := emptyBoard()
	tpl.Execute(w, TplData{BoardString: "", Board: boardArr})
}

type TplData struct {
	BoardString string
	Board       [7][6]string
}

func homePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	col := r.FormValue("Column")
	board := r.FormValue("Board")
	r.Body.Close()

	board += col

	fmt.Println("Board: ", board)

	// use engine to calculate best move given board string and save the result
	// in new variable which you append to board
	b, _ := game.MakeBoardFromString(board)
	insertCol, _ := solver.GetBestMove(b)
	board += strconv.Itoa(insertCol + 1)

	boardArr := boardFromString(emptyBoard(), board)

	tpl, err := template.ParseFiles("../../ui/html/home.gohtml")
	if err != nil {
		panic(err)
	}
	tpl.Execute(w, TplData{BoardString: board, Board: boardArr})
}
