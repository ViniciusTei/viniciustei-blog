package main

import (
	"fmt"
	"net/http"

	"github.com/ViniciusTei/viniciustei-blog/handlers"
	"github.com/ViniciusTei/viniciustei-blog/middlewares"
	"github.com/ViniciusTei/viniciustei-blog/repositories"
	"github.com/ViniciusTei/viniciustei-blog/usecases"
)

func main() {
	articleRepo := &repositories.ArticleRepositoryImpl{}
	articleUseCase := &usecases.ArticleUseCase{Repo: articleRepo}
	handler := &handlers.Handler{ArticleUseCase: articleUseCase}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", handler.HandleRoot)
	mux.HandleFunc("/article/{slug}", handler.HandleArticles)
	mux.HandleFunc("/about", handlers.HandleAbout)
	mux.HandleFunc("/login", handlers.HandleLogin)
	mux.HandleFunc("/dashboard", handlers.HandleAbout)

	//middlewares
	withMidllewaresMux := middlewares.NewLogger(mux)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", withMidllewaresMux)
}
