package render

import (
	"errors"
	"html/template"
	"net/http"
	"path/filepath"
)

var cachedTemplates = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, tmplName string) error {

	if len(cachedTemplates) == 0 {
		tmplCache, err := createTemplateCache()
		if err != nil {
			return errors.New("error creating template cache")
		}
		cachedTemplates = tmplCache

	}
	tmpl := cachedTemplates[tmplName]
	err := tmpl.Execute(w, nil)
	if err != nil {
		return errors.New("error executing template")
	}
	return nil
}

func createTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	layouts, err := filepath.Glob("./templates/*.layout.tmpl")
	if err != nil {
		return nil, errors.New("did not find layouts")
	}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return nil, errors.New("did not find files")
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tmpl, err := template.New(name).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		tmplWithLayout, err := tmpl.ParseFiles(layouts...)
		if err != nil {
			return nil, errors.New("error during parsing layouts")
		}
		templateCache[name] = tmplWithLayout

	}
	return templateCache, nil
}
