package book_repository_v1

import (
	"time"

	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"
)

type Book struct {
	Id           int        `gorm:"primary_key;column:id"`
	Title        string     `gorm:"column:title"`
	ReleasedYear int        `gorm:"column:released_year"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
}

func (b *Book) ToBookEntity(rating float64) *entity.Book {
	return &entity.Book{
		Id:           b.Id,
		Title:        b.Title,
		ReleasedYear: b.ReleasedYear,
		Rating:       rating,
	}
}
func (Book) FromBookEntity(e *entity.Book) *Book {
	return &Book{
		Id:           e.Id,
		Title:        e.Title,
		ReleasedYear: e.ReleasedYear,
	}
}

type BookWithRating struct {
	Book
	Rating float64 `gorm:"column:rating"`
}

func (b *BookWithRating) ToBookEntity() *entity.Book {
	return &entity.Book{
		Id:           b.Id,
		Title:        b.Title,
		ReleasedYear: b.ReleasedYear,
		Rating:       b.Rating,
	}
}
