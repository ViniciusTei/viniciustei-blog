package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title   string
	Content template.HTML
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

	content := "<ul>\n"
	for _, article := range articles {
		content += "<li>" + article.Title + "</li>\n"
	}
	content += "</ul>\n"

	data := PageData{
		Title:   "My Blog",
		Content: template.HTML(content),
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
