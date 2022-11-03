package main

import (
	"fmt"
	"html/template" // New import
	"io/fs"
	"path/filepath" // New import

	"as207414.net/as207414.net/web"
)

type Template struct {
	pages	map[string]*template.Template
}

func newTemplate(directory string) (*Template, error) {

	renderedPages := map[string]*template.Template{}

	pages, err := fs.Glob(web.Files, fmt.Sprintf("%s/pages/*.html", directory))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFS(web.Files, page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFS(web.Files, fmt.Sprintf("%s/layouts/*.html", directory))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFS(web.Files, fmt.Sprintf("%s/partials/*.html", directory))
		if err != nil {
			return nil, err
		}

		renderedPages[name] = ts
	}

	return &Template{ pages: renderedPages }, nil
}