package main

import (
	"as207414.net/as207414.net/web"
	"net/http"
	"io/fs"
)

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.index)
	mux.HandleFunc("/peering.html", app.peering)

	// static assets
	static, err := fs.Sub(web.Files, "static")

	if err != nil {
		panic(err)
	}

	statics := http.FileServer(http.FS(static))
	mux.Handle("/static/", statics)



	// fileServer := http.FileServer(http.FS(web.Files))
	// mux.Handle("/static/", fileServer)

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
