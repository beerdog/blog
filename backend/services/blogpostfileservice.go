package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"blog.jonastrogen.se/models"
)

type BlogpostFileService struct{}

func (s BlogpostFileService) GetBlogpost(ctx context.Context, name string) ([]byte, error) {
	content, err := os.ReadFile("blogposts/" + name + ".md")
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (s BlogpostFileService) getMetadataByKey(ctx context.Context, key string) (*models.Metadata, error) {
	fileContents, err := os.ReadFile(key)
	if err != nil {
		return nil, err
	}

	metadata := models.Metadata{}

	err = json.Unmarshal([]byte(fileContents), &metadata)
	if err != nil {
		return nil, err
	}

	return &metadata, nil
}

func (s BlogpostFileService) GetMetadata(ctx context.Context, name string) (*models.Metadata, error) {
	return s.getMetadataByKey(ctx, "blogposts/"+name+".json")
}

func (s BlogpostFileService) ListMetadata(ctx context.Context) (*[]models.Metadata, error) {
	files, err := os.ReadDir("blogposts/")
	if err != nil {
		return nil, err
	}

	metadataList := []models.Metadata{}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			metadata, err := s.getMetadataByKey(ctx, fmt.Sprintf("blogposts/%s", file.Name()))
			if err != nil {
				return nil, err
			}
			metadataList = append(metadataList, *metadata)
		}
		fmt.Println(file.Name(), file.IsDir())
	}

	return &metadataList, err
}
