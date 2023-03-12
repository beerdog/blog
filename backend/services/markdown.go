package services

import (
	"bytes"

	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

func RenderMarkdownFile(md []byte) (*bytes.Buffer, error) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
				highlighting.WithFormatOptions(
					html.WithLineNumbers(true),
				),
			),
		),
	)

	var buf bytes.Buffer
	if err := markdown.Convert([]byte(md), &buf); err != nil {
		panic(err)
	}

	return &buf, nil
}
