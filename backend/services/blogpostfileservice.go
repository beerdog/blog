package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"blog.jonastrogen.se/models"
)

type BlogpostFileService struct{}

func (s BlogpostFileService) GetMetadata(postName string) (*models.Metadata, error) {
	fileContents, err := os.ReadFile(postName)
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

func (s BlogpostFileService) ListMetadata() (*[]models.Metadata, error) {
	files, err := ioutil.ReadDir("blogposts/")
	if err != nil {
		return nil, err
	}

	// WIP fix this
	metadataList := []models.Metadata{}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			metadata, err := s.GetMetadata("blogposts/" + file.Name())
			if err != nil {
				return nil, err
			}
			metadataList = append(metadataList, *metadata)
		}
		fmt.Println(file.Name(), file.IsDir())
	}

	return &metadataList, err
}
