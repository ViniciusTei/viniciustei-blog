package tests

import (
	"os"
	"testing"

	"github.com/ViniciusTei/viniciustei-blog/internal/repositories"
)

func TestLoadArticles(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(dir+"/test1.md", []byte("# Title\nContent"), 0644)

	repo := &repositories.ArticleRepositoryImpl{}
	articles, err := repo.LoadArticles(dir)
	if err != nil {
		t.Fatalf("Failed to load articles: %v", err)
	}

	if len(articles) == 0 {
		t.Errorf("Expected at least one article, got %d", len(articles))
	}

	for _, article := range articles {
		if article.Title == "" {
			t.Error("Article title should not be empty")
		}
		if article.Slug == "" {
			t.Error("Article slug should not be empty")
		}
	}
}

func TestLoadArticleBySlug(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(dir+"/test1.md", []byte("# Title\nContent"), 0644)

	repo := &repositories.ArticleRepositoryImpl{}
	articles, err := repo.LoadArticles(dir)
	if err != nil {
		t.Fatalf("Failed to load articles: %v", err)
	}

	if len(articles) == 0 {
		t.Fatalf("Expected at least one article, got %d", len(articles))
	}

	slug := articles[0].Slug
	article, err := repo.LoadArticleBySlug(slug, dir)
	if err != nil {
		t.Fatalf("Failed to load article by slug: %v", err)
	}

	if article.Slug != slug {
		t.Errorf("Expected slug '%s', got '%s'", slug, article.Slug)
	}
}
