package server

import "blog.jonastrogen.se/models"

type BlogPostService interface {
	GetMetadata(postName string) (*models.Metadata, error)
	ListMetadata() (*[]models.Metadata, error)
}
