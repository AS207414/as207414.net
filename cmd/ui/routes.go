package main

import (
    "net/http"
    "as207414.net/as207414.net/web"
)

func (app *application) routes() http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/", app.index)
    mux.HandleFunc("/peering.html", app.peering)

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/static/", fileServer)

    return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}