package main

import (
	"net/http"
)

func (app *application) index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	app.render(w, r, "index.html")

}

func (app *application) peering(w http.ResponseWriter, r *http.Request) {

	app.render(w, r, "peering.html")

}
