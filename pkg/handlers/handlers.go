package handlers

import (
	"awesomeProject/pkg/config"
	"awesomeProject/pkg/render"
	"net/http"
)

type Repository struct {
	app *config.AppConfig
}

var Repo *Repository

func CreateRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{app: appConfig}
}

func InjectRepo(repo *Repository) {
	Repo = repo
}

// Home is the handler for the home page
func (repo *Repository) Home(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func (repo *Repository) About(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
