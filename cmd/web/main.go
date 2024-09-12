package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type application struct {
	templateCache map[string]*template.Template
}

func main() {
	mux := http.NewServeMux()

	templateCache, err := newTemplateCache()
	if err != nil {
		// TODO clean up error handling
		panic(err)
	}
	app := &application{templateCache: templateCache}
	routes := app.routes(mux)
	fmt.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", routes)
}
