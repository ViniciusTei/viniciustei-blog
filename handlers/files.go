package handlers

import (
	"os"
	"path/filepath"
	"strings"
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

	title := strings.Split(string(content), "\n")[0]
	slug := strings.Split(filepath, "/")[len(strings.Split(filepath, "/"))-1]
	return Article{Title: strings.TrimPrefix(title, "# "), Content: string(content), Slug: slug}, nil
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
