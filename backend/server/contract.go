package server

import (
	"context"

	"blog.jonastrogen.se/models"
)

type BlogPostService interface {
	GetBlogpost(ctx context.Context, name string) ([]byte, error)
	GetMetadata(ctx context.Context, postName string) (*models.Metadata, error)
	ListMetadata(ctx context.Context) (*[]models.Metadata, error)
}
