package main

import (
    "net/http"
	"html/template"
)

func (app *application) index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

	files := []string{
		"./web/templates/index.page.html",
		"./web/templates/base.layout.html",
		"./web/templates/footer.partial.html",
		"./web/templates/header.partial.html",

	}

	ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

	err = ts.Execute(w, nil)
    if err != nil {
        app.serverError(w, err)
    }

}

func (app *application) peering(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./web/templates/peering.page.html",
		"./web/templates/base.layout.html",
		"./web/templates/footer.partial.html",
		"./web/templates/header.partial.html",

	}

	ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, err)
        return
    }

	err = ts.Execute(w, nil)
    if err != nil {
        app.serverError(w, err)
    }


}