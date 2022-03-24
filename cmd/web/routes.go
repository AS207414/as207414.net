package main

import "net/http"

func (app *application) routes() http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/", app.index)
    mux.HandleFunc("/peering.html", app.peering)

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}