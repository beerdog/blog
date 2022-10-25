package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetMetadata(file string) (*Metadata, error) {
	fileContents, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	metadata := Metadata{}

	err = json.Unmarshal([]byte(fileContents), &metadata)
	if err != nil {
		return nil, err
	}

	return &metadata, nil
}

func ListMetadata() (*[]Metadata, error) {
	files, err := ioutil.ReadDir("blogposts/")
	if err != nil {
		return nil, err
	}

	// WIP fix this
	metadataList := []Metadata{}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			metadata, err := GetMetadata("blogposts/" + file.Name())
			if err != nil {
				return nil, err
			}
			metadataList = append(metadataList, *metadata)
		}
		fmt.Println(file.Name(), file.IsDir())
	}

	return &metadataList, err
}
