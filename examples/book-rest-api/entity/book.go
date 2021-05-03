package entity

import (
	"time"
)

type Book struct {
	Id           int          `json:"id"`
	Title        string       `json:"title"`
	ReleasedYear int          `json:"released_year"`
	CreatedAt    *time.Time   `json:"created_at"`
	UpdatedAt    *time.Time   `json:"updated_at"`
	Authors      []BookAuthor `json:"authors"`
}

func (b *Book) AddAuthor(ba BookAuthor) {
	if b.Authors == nil {
		b.Authors = []BookAuthor{}
	}
	b.Authors = append(b.Authors, ba)
}

type BookAuthor struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
