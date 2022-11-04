package main

import (
	"as207414.net/as207414.net/web"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"io/fs"
)

func (app *application) routes() http.Handler {

	mux := httprouter.New()
	mux.GET("/", app.index)
	mux.GET("/peering.html", app.peering)

	// static assets
	static, err := fs.Sub(web.Files, "static")

	if err != nil {
		panic(err)
	}

	mux.ServeFiles("/static/*filepath", http.FS(static))

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
