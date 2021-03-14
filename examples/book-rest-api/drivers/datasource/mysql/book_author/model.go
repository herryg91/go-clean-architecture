package book_author_datasource

import (
	"time"
)

type BookAuthorModel struct {
	BookId   int `gorm:"column:book_id"`
	AuthorId int `gorm:"column:author_id"`

	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}
