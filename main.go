package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ViniciusTei/viniciustei-blog/handlers"
	"github.com/ViniciusTei/viniciustei-blog/repositories"
	"github.com/ViniciusTei/viniciustei-blog/usecases"
)

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	currentTime, err := time.Parse(time.RFC822, time.Now().Format(time.RFC822))
	if err != nil {
		panic(err)
	}

	fmt.Printf("[%s]: %s %s %v\n", currentTime, r.Method, r.URL.Path, time.Since(start))
}

func NewLogger(handler http.Handler) *Logger {
	return &Logger{handler: handler}
}

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

	//middlewares
	loggerMux := NewLogger(mux)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", loggerMux)
}
