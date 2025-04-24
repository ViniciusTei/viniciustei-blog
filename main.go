package main

import (
	"fmt"
	"net/http"

	"github.com/ViniciusTei/viniciustei-blog/handlers"
)

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", handlers.HandleRoot)
	mux.HandleFunc("/article/{slug}", handlers.HandleArticles)
	mux.HandleFunc("/about", handlers.HandleAbout)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}
