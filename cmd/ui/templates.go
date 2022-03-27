package main

import (
	"html/template" // New import
	"io/fs"
	"path/filepath" // New import

	"as207414.net/as207414.net/web"
)

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(web.Files, "templates/*.page.html")
	if err != nil {
		return nil, err
	}

	// Loop through the pages one-by-one.
	for _, page := range pages {
		name := filepath.Base(page)

		// Parse the page template file in to a template set.
		ts, err := template.New(name).ParseFS(web.Files, page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFS(web.Files, "templates/*.layout.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFS(web.Files, "templates/*.partial.html")
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	// Return the map.
	return cache, nil
}
