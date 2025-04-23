package main

import (
	"fmt"
	"net/http"

	"github.com/ViniciusTei/viniciustei-blog/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HandleRoot)
	mux.HandleFunc("/articles", handlers.HandleArticles)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}
