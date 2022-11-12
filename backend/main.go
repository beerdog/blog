package main

import (
	"context"
	"fmt"
	"net/http"

	"blog.jonastrogen.se/server"
	"blog.jonastrogen.se/services"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Is set in makefile when building for serverless deployment.
var Environment = "local"

var r *chi.Mux

func Handler(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	httpReq, err := server.EventToRequest(req)
	if err != nil {
		return events.LambdaFunctionURLResponse{
			Headers:    map[string]string{},
			StatusCode: http.StatusInternalServerError,
		}, err // FIXME do not expose errors in prod
	}
	respWriter := server.NewProxyResponseWriter()
	r.ServeHTTP(respWriter, httpReq)
	proxyResponse, err := respWriter.GetProxyResponse()
	if err != nil {
		return events.LambdaFunctionURLResponse{
			Headers:    map[string]string{},
			StatusCode: http.StatusInternalServerError,
		}, err // FIXME do not expose errors in prod
	}

	return proxyResponse, nil
}

func main() {
	if Environment == "local" {
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

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./templates/index.html")
	// })
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	awsConfig, err := awsconfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(fmt.Errorf("failed to load default AWS config: %w", err))
	}

	var blogpostService server.BlogPostService
	if Environment == "local" {
		//blogpostService = services.BlogpostFileService{}
		//blogpostService = services.NewBlogpostS3Service("trogen", awsConfig)
		blogpostService = services.NewBlogpostDynamoDBService("blogposts", awsConfig)
	} else {
		// blogpostService = services.NewBlogpostS3Service("trogen", awsConfig)
		blogpostService = services.NewBlogpostDynamoDBService("blogposts", awsConfig)
	}

	h := server.NewHandler(blogpostService)

	r.MethodFunc("get", "/api/blogposts/{blogpost}", h.HandleGetBlogpost)
	r.MethodFunc("get", "/api/metadata/{blogpost}", h.HandleGetMetadata)
	r.MethodFunc("get", "/api/metadata", h.HandleListMetadata)
	r.MethodFunc("get", "/api/test", h.HandleTest)

	fmt.Printf("Stage: %s\n", Environment)
}
