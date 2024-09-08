package main

import "net/http"

func routes(mux *http.ServeMux) http.Handler {
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../ui/static/css"))))
	mux.HandleFunc("GET /", home)
	mux.HandleFunc("POST /", homePost)
	return mux
}
