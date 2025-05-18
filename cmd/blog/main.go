package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	db "github.com/ViniciusTei/viniciustei-blog/internal/database"
	"github.com/ViniciusTei/viniciustei-blog/internal/handlers"
	"github.com/ViniciusTei/viniciustei-blog/internal/middlewares"
	"github.com/ViniciusTei/viniciustei-blog/internal/repositories"
	"github.com/ViniciusTei/viniciustei-blog/utils"
	mux "github.com/gorilla/mux"
)

func main() {
	config := utils.LoadConfig()
	database, err := db.Conn(config.DBUrl)
	if err != nil {
		panic(err)
	}

	articleRepo := repositories.NewArticleRepository(&database)

	handler := &handlers.Handler{
		ArticleUseCase: articleRepo,
	}

	r := mux.NewRouter()

	//middlewares
	r.Use(middlewares.NewLogger().ServeHTTP)
	r.Use(middlewares.NewResponseHeader("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0").ServeHTTP)
	r.Use(mux.CORSMethodMiddleware(r))

	r.HandleFunc("/", handler.HandleRoot)
	r.HandleFunc("/about", handler.HandleAbout)
	r.HandleFunc("/login", handler.HandleLogin)

	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	app := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("localhost:%s", config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Starting server on %s\n", app.Addr)

	log.Fatal(app.ListenAndServe())
}
