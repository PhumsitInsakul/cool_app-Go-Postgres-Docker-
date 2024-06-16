package app

import (
	"html/template"
	"net/http"
)

// Structure for store templates in memory
type Templates struct {
	templates *template.Template
}

func NewTemplates() *Templates {
	// return a new instance of the Templates struct
	return &Templates{
		// read all templates from the templates folder
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}

func (t *Templates) Render(w http.ResponseWriter, name string, data map[string]any) {
	t.templates.ExecuteTemplate(w, name, data)
}
