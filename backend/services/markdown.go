package services

import (
	"bytes"
	"os"

	"github.com/yuin/goldmark"
)

func RenderMarkdownFile(file string) (*bytes.Buffer, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(content, &buf); err != nil {
		return nil, err
	}

	return &buf, nil
}
