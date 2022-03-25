package main

import (
	"fmt"
	"net/http"
)

func (app *application) serve() error {

	srv := &http.Server{
		Addr:		fmt.Sprintf("%s:%d",app.config.address, app.config.port),
		ErrorLog: 	app.errorLog,
		Handler: 	app.routes(),
	}

	// Start the server as normal, returning any error.
    return srv.ListenAndServe()
}