package repositories

import (
	"context"
	"log"

	"github.com/ViniciusTei/viniciustei-blog/internal/database"
	"github.com/ViniciusTei/viniciustei-blog/internal/entities"
)

type ArticleRepositoryImpl struct {
	db *database.DatabaseImpl
}

func NewArticleRepository(db *database.DatabaseImpl) *ArticleRepositoryImpl {
	return &ArticleRepositoryImpl{
		db: db,
	}
}

func (r *ArticleRepositoryImpl) LoadArticles() ([]entities.Article, error) {
	rows, err := r.
		db.
		Conn.
		Query(context.Background(), "SELECT * FROM artigos")
	if err != nil {
		//TODO: handle SQL error and return a more user-friendly error
		log.Printf("Database query error: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var articles []entities.ArticleModel

	for rows.Next() {
		var a entities.ArticleModel
		err := rows.Scan(&a.Id, &a.Title, &a.Content, &a.CriadoEm, &a.AtualizadoEm, &a.AuthorId, &a.Slug)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		articles = append(articles, a)
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
			Date:        a.CriadoEm,
			Description: "This is a sample description for the article.",
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
