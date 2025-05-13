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
	"github.com/ViniciusTei/viniciustei-blog/utils"
)

//go:embed templates/*.html
var templatesFS embed.FS

func main() {
	config := utils.LoadConfig()
	database, err := db.Conn(config.DBUrl)
	if err != nil {
		panic(err)
	}

	// Repositories
	articleRepo := &repositories.ArticleRepositoryImpl{Db: &database}
	authRepo := &repositories.AuthRepositoryImpl{Db: &database}

	// Use Cases
	articleUseCase := &usecases.ArticleUseCase{Repo: articleRepo}
	authUseCase := &usecases.AuthUseCase{Repo: authRepo}

	handler := &handlers.Handler{
		ArticleUseCase: articleUseCase,
		AuthUseCase:    authUseCase,
		Views:          templatesFS,
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("GET /article/{slug}", handler.HandleArticles)
	mux.HandleFunc("GET /about", handler.HandleAbout)
	mux.HandleFunc("GET /login", handler.HandleLogin)
	mux.HandleFunc("POST /signin", handler.HandleSignIn)
	mux.HandleFunc("POST /signout", handlers.HandleSignOut)
	mux.HandleFunc("/", handler.HandleRoot)

	//middlewares
	withMidllewaresMux := middlewares.NewLogger(middlewares.NewResponseHeader(mux, "Cache-Control", "no-store, no-cache, must-revalidate, max-age=0"))

	fmt.Printf("Starting server on %s\n", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), withMidllewaresMux)
}
