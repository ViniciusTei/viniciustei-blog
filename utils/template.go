package utils

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, page string, data interface{}) {
	tmpl, err := template.ParseFiles(
		filepath.Join("static", "templates", "layout.html"),
		filepath.Join("static", "templates", page),
	)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
