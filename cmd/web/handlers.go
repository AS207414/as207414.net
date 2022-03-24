package main

import (
    "net/http"
	"log"
	"html/template"
)

func index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

	files := []string{
		"./web/templates/index.page.html",
		"./web/templates/base.layout.html",
		"./web/templates/footer.partial.html",
		"./web/templates/header.partial.html",

	}

	ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

	err = ts.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }

}

func peering(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./web/templates/peering.page.html",
		"./web/templates/base.layout.html",
		"./web/templates/footer.partial.html",
		"./web/templates/header.partial.html",

	}

	ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

	err = ts.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }


}