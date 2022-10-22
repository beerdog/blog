package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// HandleGet serves a HTML page
// func (h *Handlers) HandleGet(w http.ResponseWriter, r *http.Request) {
// 	t, _ := template.ParseFiles("templates/index.gohtml")
// 	t.Execute(w)
// }

func HandleGet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func HandleGetArticle(w http.ResponseWriter, r *http.Request) {
	blogpost := chi.URLParam(r, "blogpost")
	markdown, err := RenderMarkdownFile("blogposts/" + blogpost + ".md")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(markdown.Bytes())
}

func HandleGetArticleMetadata(w http.ResponseWriter, r *http.Request) {
	blogpost := chi.URLParam(r, "blogpost")
	metadata, err := GetMetadata("blogposts/" + blogpost + ".json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	SendJSONResponse(w, metadata)
}
