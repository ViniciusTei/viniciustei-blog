package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ViniciusTei/viniciustei-blog/internal/repositories"
	"github.com/gorilla/mux"
)

type ArticlePage struct {
	Title   string
	Content template.HTML
}

type ArticleController struct {
	article *repositories.ArticleRepositoryImpl
}

func NewArticleController(articleUseCase *repositories.ArticleRepositoryImpl) *ArticleController {
	return &ArticleController{
		article: articleUseCase,
	}
}

func (ac *ArticleController) Pages(mux *mux.Router) {
	mux.HandleFunc("GET /article/{slug}", ac.handleArticles)
}

func (h *ArticleController) handleArticles(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	article, err := h.article.LoadArticleBySlug(slug)
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
