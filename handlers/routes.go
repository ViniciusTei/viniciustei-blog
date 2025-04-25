package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ViniciusTei/viniciustei-blog/entities"
	"github.com/ViniciusTei/viniciustei-blog/usecases"
)

type PageData struct {
	Title    string
	Articles []entities.Article
	UserId   string
	UserName string
}

type Handler struct {
	ArticleUseCase *usecases.ArticleUseCase
}

func (h *Handler) HandleRoot(w http.ResponseWriter, r *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	articles, err := h.ArticleUseCase.GetAllArticles()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error loading articles:", err)
		return
	}

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	data := PageData{
		Title:    "My Blog",
		Articles: articles,
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		return
	}
}

type ArticlePage struct {
	Title    string
	Content  template.HTML
	UserId   string
	UserName string
}

func (h *Handler) HandleArticles(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	article, err := h.ArticleUseCase.GetArticleBySlug(slug)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error getting article:", err)
		return
	}

	if article.Slug == "" {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/articles.html",
	)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", ArticlePage{
		Title:   article.Title,
		Content: template.HTML(article.Content),
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		return
	}
}

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/about.html",
	)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	data := PageData{
		Title: "About",
		// Adicione outros campos se necessário
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		return
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/login.html",
	)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	data := PageData{
		Title: "Login",
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		return
	}
}
