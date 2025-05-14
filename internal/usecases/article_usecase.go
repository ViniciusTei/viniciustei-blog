package usecases

import (
	"github.com/ViniciusTei/viniciustei-blog/internal/entities"
)

type ArticleRepository interface {
	LoadArticles() ([]entities.Article, error)
	LoadArticleBySlug(slug string) (entities.Article, error)
}

type ArticleUseCase struct {
	Repo ArticleRepository
}

func (uc *ArticleUseCase) GetAllArticles() ([]entities.Article, error) {
	return uc.Repo.LoadArticles()
}

func (uc *ArticleUseCase) GetArticleBySlug(slug string) (entities.Article, error) {
	return uc.Repo.LoadArticleBySlug(slug)
}
