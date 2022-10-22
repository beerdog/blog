package main

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

	// err = os.WriteFile("test.html", buf.Bytes(), 0644)
	// if err != nil {
	// 	return nil, err
	// }
	return &buf, nil
}
