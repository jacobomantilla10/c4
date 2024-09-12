package main

import "html/template"

type templateData struct {
	BoardString string
	Board       [7][6]string
	IsGameOver  bool
	Outcome     string
}

func newTemplateCache() (map[string]*template.Template, error) {
	templateCache := make(map[string]*template.Template)

	tpl, err := template.ParseFiles("../../ui/html/home.html", "../../ui/html/partials/board.html")
	if err != nil {
		return nil, err
	}
	templateCache["home"] = tpl

	tpl, err = template.ParseFiles("../../ui/html/partials/board.html")
	if err != nil {
		return nil, err
	}

	templateCache["board"] = tpl

	return templateCache, nil
}

func newTemplateData() templateData {
	return templateData{
		BoardString: "",
		Board:       emptyBoard(),
		IsGameOver:  false,
		Outcome:     "Game will end in a draw",
	}
}
