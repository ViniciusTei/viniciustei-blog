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
