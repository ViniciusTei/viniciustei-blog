package handlers

import (
	"log"
	"net/http"

	"github.com/ViniciusTei/viniciustei-blog/internal/entities"
	"github.com/ViniciusTei/viniciustei-blog/internal/repositories"
	"github.com/ViniciusTei/viniciustei-blog/utils"
)

type PageData struct {
	Title    string
	Articles []entities.Article
	UserId   string
	UserName string
	Error    string
}

type Handler struct {
	ArticleUseCase *repositories.ArticleRepositoryImpl
}

func (h *Handler) HandleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	articles, err := h.ArticleUseCase.LoadArticles()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error loading articles:", err)
		return
	}

	userId, userName, err := utils.GetUserFromCookie(r)
	if err == nil {
		log.Println("Error getting cookies:", err) // yet do nothing
	}

	data := PageData{
		Title:    "My Blog",
		Articles: articles,
		UserId:   userId,
		UserName: userName,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	utils.RenderTemplate(w, "index.html", data)
}

func (h *Handler) HandleAbout(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "About",
		// Adicione outros campos se necessário
	}

	utils.RenderTemplate(w, "about.html", data)
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	err := r.URL.Query().Get("error")
	data := PageData{
		Title: "Login",
		Error: err,
	}

	utils.RenderTemplate(w, "login.html", data)
}
