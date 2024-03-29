package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"blog.jonastrogen.se/models"
	"blog.jonastrogen.se/services"
)

type BlogPostHandler struct {
	BlogPostService BlogPostService
}

func NewHandler(blogpostService BlogPostService) *BlogPostHandler {
	return &BlogPostHandler{
		BlogPostService: blogpostService,
	}
}

func (h *BlogPostHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func (h *BlogPostHandler) HandleGetBlogpost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	blogpost := chi.URLParam(r, "blogpost")
	md, err := h.BlogPostService.GetBlogpost(r.Context(), blogpost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	markdown, err := services.RenderMarkdownFile([]byte(md.Content))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(markdown.Bytes())
}

func (h *BlogPostHandler) HandleGetMetadata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	blogpost := chi.URLParam(r, "blogpost")
	metadata, err := h.BlogPostService.GetMetadata(r.Context(), blogpost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	SendJSONResponse(w, metadata)
}

func (h *BlogPostHandler) HandleListMetadata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	metadata, err := h.BlogPostService.ListMetadata(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	SendJSONResponse(w, metadata)
}

func (h *BlogPostHandler) HandleTest(w http.ResponseWriter, r *http.Request) {
	metadata := models.Metadata{
		Title:       "Test",
		Preamble:    "Lorem ipsum dolaret",
		PublishDate: models.NewDate(time.Now()),
		Category:    "Software",
		Tags:        []string{"aws", "lambda", "go", "webhook", "serverless"},
	}
	SendJSONResponse(w, metadata)
}
