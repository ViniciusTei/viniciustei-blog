package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title    string
	Articles []Article
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	articles, err := LoadMarkdownFiles("articles")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error loading articles:", err)
		return
	}

	data := PageData{
		Title:    "My Blog",
		Articles: articles,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		return
	}

}

func HandleArticles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Articles Page"))
}
