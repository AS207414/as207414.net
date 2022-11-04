package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (app *application) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	app.render(w, r, "index.html")

}

func (app *application) peering(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	app.render(w, r, "peering.html")

}
