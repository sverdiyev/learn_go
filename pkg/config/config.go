package config

import "html/template"

type AppConfig struct {
	Templates map[string]*template.Template
}
