package services

import (
	"bytes"

	"github.com/yuin/goldmark"
)

func RenderMarkdownFile(md []byte) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert(md, &buf); err != nil {
		return nil, err
	}

	return &buf, nil
}
