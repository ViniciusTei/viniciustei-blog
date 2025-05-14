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

	articleRepo := repositories.NewArticleRepository(&database, "cmd/blog/static/articles")
	authRepo := &repositories.AuthRepositoryImpl{Db: &database}

	articleUseCase := &usecases.ArticleUseCase{Repo: articleRepo}
	authUseCase := &usecases.AuthUseCase{Repo: authRepo}

	uc := handlers.NewUserController(authUseCase)
	ac := handlers.NewArticleController(articleUseCase, templatesFS)

	handler := &handlers.Handler{
		ArticleUseCase: articleUseCase,
		Views:          templatesFS,
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	uc.Pages(mux)
	uc.Routes(mux)
	ac.Pages(mux)

	mux.HandleFunc("GET /about", handler.HandleAbout)
	mux.HandleFunc("GET /login", handler.HandleLogin)
	mux.HandleFunc("/", handler.HandleRoot)

	//middlewares
	withMidllewaresMux := middlewares.NewLogger(middlewares.NewResponseHeader(mux, "Cache-Control", "no-store, no-cache, must-revalidate, max-age=0"))

	fmt.Printf("Starting server on %s\n", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), withMidllewaresMux)
}
