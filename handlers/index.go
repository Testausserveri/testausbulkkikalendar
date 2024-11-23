package handlers

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func Init(templateGlob string) {
	templates = template.Must(template.ParseGlob(templateGlob))
}

func Index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Testausbulkkikalendar",
		Message: "WIP",
	}

	// Render the "index.html" template
	templates.ExecuteTemplate(w, "index.html", data)
}
