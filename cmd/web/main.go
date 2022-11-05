package main

import (
	"awesomeProject/pkg/config"
	"awesomeProject/pkg/handlers"
	"awesomeProject/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	cachedTmpls, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln(err)
	}
	render.InjectCachedTemplates(&app)

	app.Templates = cachedTmpls

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
