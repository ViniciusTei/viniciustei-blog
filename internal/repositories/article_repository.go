package repositories

import (
	"time"

	"github.com/ViniciusTei/viniciustei-blog/internal/entities"
	"github.com/ViniciusTei/viniciustei-blog/internal/handlers"
)

type ArticleRepositoryImpl struct{}

// TODO: implement a database to store article metadata
func (r *ArticleRepositoryImpl) LoadArticles(dir string) ([]entities.Article, error) {
	// Load articles using the existing function
	articles, err := handlers.LoadMarkdownFiles(dir)
	if err != nil {
		return nil, err
	}

	// Convert handlers.Article to entities.Article
	var result []entities.Article
	for _, a := range articles {
		result = append(result, entities.Article{
			Title:       a.Title,
			Content:     a.Content,
			Slug:        a.Slug,
			Image:       "https://placehold.co/600x400/png",
			Category:    "General",
			ReadTime:    "5 min",
			Date:        time.Now(),
			Description: "This is a sample description for the article.",
		})
	}

	return result, nil
}

func (r *ArticleRepositoryImpl) LoadArticleBySlug(slug string, dir string) (entities.Article, error) {
	articles, err := r.LoadArticles(dir)
	if err != nil {
		return entities.Article{}, err
	}

	for _, a := range articles {
		if a.Slug == slug {
			return a, nil
		}
	}

	return entities.Article{}, nil
}
