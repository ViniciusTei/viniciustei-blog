package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ViniciusTei/viniciustei-blog/internal/usecases"
)

type ArticleController struct {
	articleUseCase *usecases.ArticleUseCase
}

func NewArticleController(articleUseCase *usecases.ArticleUseCase) *ArticleController {
	return &ArticleController{
		articleUseCase: articleUseCase,
	}
}

func (ac *ArticleController) Pages(mux *http.ServeMux) {
	mux.HandleFunc("GET /article/{slug}", ac.handleArticles)
}

func (h *ArticleController) handleArticles(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	article, err := h.articleUseCase.GetArticleBySlug(slug)
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
