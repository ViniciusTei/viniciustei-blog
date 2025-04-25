package usecases

import "github.com/ViniciusTei/viniciustei-blog/entities"

type ArticleRepository interface {
	LoadArticles(dir string) ([]entities.Article, error)
	LoadArticleBySlug(slug string, dir string) (entities.Article, error)
}

type ArticleUseCase struct {
	Repo ArticleRepository
}

func (uc *ArticleUseCase) GetAllArticles() ([]entities.Article, error) {
	return uc.Repo.LoadArticles("articles")
}

func (uc *ArticleUseCase) GetArticleBySlug(slug string) (entities.Article, error) {
	return uc.Repo.LoadArticleBySlug(slug, "articles")
}
