package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"

	"blog.jonastrogen.se/server"
	"blog.jonastrogen.se/services"
)

var chiLambda *chiadapter.ChiLambda

// handler is the function called by the lambda.
func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return chiLambda.ProxyWithContext(ctx, req)
}

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

	blogpostService := services.BlogpostFileService{}

	h := server.NewHandler(blogpostService)

	r.MethodFunc("get", "/api/blogposts/{blogpost}", h.HandleGetBlogpost)
	r.MethodFunc("get", "/api/metadata/{blogpost}", h.HandleGetMetadata)
	r.MethodFunc("get", "/api/metadata", h.HandleListMetadata)

	// start the lambda with a context
	lambda.StartWithOptions(handler, lambda.WithContext(context.TODO()))
}
