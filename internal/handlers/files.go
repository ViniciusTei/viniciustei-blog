package handlers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ViniciusTei/viniciustei-blog/internal/markdown"
)

type Article struct {
	Title   string
	Content string
	Slug    string
}

func LoadMarkdown(filepath string) (Article, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return Article{}, err
	}

	title := strings.TrimPrefix(strings.Split(string(content), "\n")[0], "# ")
	slug := strings.ToLower(strings.ReplaceAll(title, " ", "-"))
	htmlContent := markdown.MarkdownToHTML(string(content))
	return Article{Title: title, Content: htmlContent, Slug: slug}, nil
}

func LoadMarkdownFiles(dir string) ([]Article, error) {
	var articles []Article

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			filePath := filepath.Join(dir, file.Name())
			article, err := LoadMarkdown(filePath)
			if err != nil {
				return nil, err
			}

			articles = append(articles, article)
		}
	}

	return articles, nil
}
