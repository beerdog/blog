package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	chi.NewMux()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/index.html")
	})

	r.MethodFunc("get", "/api/blogposts/{blogpost}", HandleGetBlogpost)
	r.MethodFunc("get", "/api/metadata/{blogpost}", HandleGetMetadata)
	r.MethodFunc("get", "/api/metadata", HandleListMetadata)

	http.ListenAndServe(":3000", r)
}
