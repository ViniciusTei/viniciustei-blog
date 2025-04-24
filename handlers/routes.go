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

type ArticlePage struct {
	Title   string
	Content template.HTML
}

func HandleArticles(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	articles, err := LoadMarkdownFiles("articles")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error loading articles:", err)
		return
	}

	var article Article
	for _, a := range articles {
		if a.Slug == slug {
			article = a
			break
		}
	}

	if article.Slug == "" {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/articles.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	err = tmpl.Execute(w, ArticlePage{
		Title:   article.Title,
		Content: template.HTML(article.Content),
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		return
	}
}
