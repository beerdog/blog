package main

import (
	"context"
	"fmt"
	"net/http"

	"blog.jonastrogen.se/server"
	"blog.jonastrogen.se/services"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Is set in makefile when building for serverless deployment.
var Stage = "local"

var r *chi.Mux

func Handler(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	httpReq, err := server.EventToRequest(req)
	if err != nil {
		return events.LambdaFunctionURLResponse{
			Headers:    map[string]string{},
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	respWriter := server.NewProxyResponseWriter()
	r.ServeHTTP(respWriter, httpReq)
	proxyResponse, err := respWriter.GetProxyResponse()
	if err != nil {
		return events.LambdaFunctionURLResponse{
			Headers:    map[string]string{},
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return proxyResponse, nil
}

func main() {
	if Stage == "local" {
		http.ListenAndServe(":3000", r)
	} else {
		lambda.Start(Handler)
	}
}

func init() {
	r = chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/index.html")
	})
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })

	blogpostService := services.BlogpostFileService{}

	h := server.NewHandler(blogpostService)

	r.MethodFunc("get", "/api/blogposts/{blogpost}", h.HandleGetBlogpost)
	r.MethodFunc("get", "/api/metadata/{blogpost}", h.HandleGetMetadata)
	r.MethodFunc("get", "/api/metadata", h.HandleListMetadata)
	r.MethodFunc("get", "/api/test", h.HandleTest)

	fmt.Printf("Stage: %s\n", Stage)
}
