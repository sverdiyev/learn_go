package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var cachedTemplates = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, tmplName string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := cachedTemplates[tmplName]
	if !inMap {
		// need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(tmplName)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in the cache
		log.Println("using cached template")
	}

	tmpl = cachedTemplates[tmplName]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(tmplName string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", tmplName),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache (map)
	cachedTemplates[tmplName] = tmpl

	return nil
}
