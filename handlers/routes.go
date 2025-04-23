package handlers

import "net/http"

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func HandleArticles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Articles Page"))
}
