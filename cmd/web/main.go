package main

import (
	"log"
	"net/http"
)


func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/peering.html", peering)

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}