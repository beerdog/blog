package main

import (
	"encoding/json"
	"os"
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
