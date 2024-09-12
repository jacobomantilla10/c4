package main

import "net/http"

func (app *application) routes(mux *http.ServeMux) http.Handler {
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../ui/static/css"))))
	mux.HandleFunc("GET /", app.home)
	mux.HandleFunc("POST /restart", app.restart)
	mux.HandleFunc("GET /game/", app.getNextMove)
	return mux
}
