package entities

import "time"

type Article struct {
	Title       string
	Content     string
	Description string
	Slug        string
	Image       string
	Category    string
	ReadTime    string
	Date        time.Time
}

// represents the article model in the database
type ArticleModel struct {
	Id           int
	Title        string
	Content      string
	CriadoEm     time.Time
	AtualizadoEm time.Time
	AuthorId     int
	Slug         string
}
