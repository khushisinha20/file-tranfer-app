package handlers

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r* http.Request) {
	renderTemplate(w, "./templates/index.html", nil)
}

func renderTemplate(w http.ResponseWriter, filePath string, data interface{}) {
	t, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, "Error parsing template: " + err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Error parsing template: " + err.Error(), http.StatusInternalServerError)
	}
}