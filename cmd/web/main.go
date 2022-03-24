package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)


func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.LUTC|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.LUTC|log.Ltime|log.Llongfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/peering.html", peering)

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog.Printf("Starting server on %s", *addr)
    err := http.ListenAndServe(*addr, mux)
    errorLog.Fatal(err)
}