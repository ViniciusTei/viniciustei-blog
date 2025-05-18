package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/ViniciusTei/viniciustei-blog/internal/entities"
	"github.com/ViniciusTei/viniciustei-blog/internal/repositories"
	"github.com/ViniciusTei/viniciustei-blog/utils"
)

type PageData struct {
	Title    string
	Articles []entities.Article
	UserId   string
	UserName string
}

type Handler struct {
	ArticleUseCase *repositories.ArticleRepositoryImpl
}

func (h *Handler) HandleRoot(w http.ResponseWriter, r *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	log.Println("Root page", r.URL.Path)
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

	var userId, userName string
	cookie, err := r.Cookie("auth_token")
	if err == nil {
		log.Println("Error getting cookies:", err) // yet do nothing
	}

	if cookie != nil && cookie.Value != "" {
		key := []byte("6091835705053067")
		decrypted, err := utils.Decrypt(key, cookie.Value)
		if err != nil {
			log.Println("Error decrypting token:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		userId = strings.Split(decrypted, ":")[0]
		userName = strings.Split(decrypted, ":")[1]
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
	log.Println("Login page")
	data := PageData{
		Title: "Login",
	}

	utils.RenderTemplate(w, "login.html", data)
}
