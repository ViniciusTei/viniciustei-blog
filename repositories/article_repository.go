package repositories

import (
	"github.com/ViniciusTei/viniciustei-blog/entities"
	"github.com/ViniciusTei/viniciustei-blog/handlers"
)

type ArticleRepositoryImpl struct{}

func (r *ArticleRepositoryImpl) LoadArticles() ([]entities.Article, error) {
	// Load articles using the existing function
	articles, err := handlers.LoadMarkdownFiles("articles")
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

func (r *ArticleRepositoryImpl) LoadArticleBySlug(slug string) (entities.Article, error) {
	articles, err := r.LoadArticles()
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
