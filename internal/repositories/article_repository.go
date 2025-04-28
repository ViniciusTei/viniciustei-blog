package repositories

import (
	"github.com/ViniciusTei/viniciustei-blog/internal/entities"
	"github.com/ViniciusTei/viniciustei-blog/internal/handlers"
)

type ArticleRepositoryImpl struct{}

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
			Title:   a.Title,
			Content: a.Content,
			Slug:    a.Slug,
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
