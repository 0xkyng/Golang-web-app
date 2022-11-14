package config

import "text/template"


// AppConfig holds application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}