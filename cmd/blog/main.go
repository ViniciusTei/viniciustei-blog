package main

import (
	"embed"
	"fmt"
	"net/http"

	db "github.com/ViniciusTei/viniciustei-blog/internal/database"
	"github.com/ViniciusTei/viniciustei-blog/internal/handlers"
	"github.com/ViniciusTei/viniciustei-blog/internal/middlewares"
	"github.com/ViniciusTei/viniciustei-blog/internal/repositories"
	"github.com/ViniciusTei/viniciustei-blog/internal/usecases"
)

//go:embed templates/*.html
var templatesFS embed.FS

func main() {
	database, err := db.Conn()
	if err != nil {
		panic(err)
	}

	// Repositories
	articleRepo := &repositories.ArticleRepositoryImpl{Db: database}
	authRepo := &repositories.AuthRepositoryImpl{Db: database}

	// Use Cases
	articleUseCase := &usecases.ArticleUseCase{Repo: articleRepo}
	authUseCase := &usecases.AuthUseCase{Repo: authRepo}

	handler := &handlers.Handler{
		ArticleUseCase: articleUseCase,
		AuthUseCase:    authUseCase,
		Templates:      templatesFS,
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/article/{slug}", handler.HandleArticles)
	mux.HandleFunc("/about", handler.HandleAbout)
	mux.HandleFunc("/login", handler.HandleLogin)
	mux.HandleFunc("POST /signin", handler.HandleSignIn)
	mux.HandleFunc("/signout", handlers.HandleSignOut)
	mux.HandleFunc("/", handler.HandleRoot)

	//middlewares
	withMidllewaresMux := middlewares.NewLogger(middlewares.NewResponseHeader(mux, "Cache-Control", "no-store, no-cache, must-revalidate, max-age=0"))

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", withMidllewaresMux)
}
