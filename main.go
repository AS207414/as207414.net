package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var index = template.Must(template.ParseFiles("web/templates/index.html"))
var error_page = template.Must(template.ParseFiles("web/templates/500.html"))
var not_found_page = template.Must(template.ParseFiles("web/templates/404.html"))
var peering = template.Must(template.ParseFiles("web/templates/peering.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	error_page.Execute(w, nil)
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	not_found_page.Execute(w, nil)
}

func peeringHandler(w http.ResponseWriter, r *http.Request) {
	peering.Execute(w, nil)
}


func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	// Assets handler
	assets_fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", assets_fs))

	// Index page
	mux.HandleFunc("/", indexHandler)

	// Peering page
	mux.HandleFunc("/peering.html", peeringHandler)

	// 404 Page
	mux.HandleFunc("/404.html", notfoundHandler)

	// 500 Page
	mux.HandleFunc("/500.html", errorHandler)

	if err := http.ListenAndServe(":" + port, mux); err != nil {
		log.Fatal(err)
	}
}