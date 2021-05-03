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

func (b *Book) ToBookEntity() *entity.Book {
	return &entity.Book{
		Id:           b.Id,
		Title:        b.Title,
		ReleasedYear: b.ReleasedYear,
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
	}
}
func (Book) FromBookEntity(e *entity.Book) *Book {
	return &Book{
		Id:           e.Id,
		Title:        e.Title,
		ReleasedYear: e.ReleasedYear,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
	}
}

type BookAuthor struct {
	BookId   int `gorm:"column:book_id"`
	AuthorId int `gorm:"column:author_id"`

	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (ba *BookAuthor) ToBookAuthorEntity() *entity.BookAuthor {
	return &entity.BookAuthor{
		Id:   ba.AuthorId,
		Name: "",
	}
}

func (BookAuthor) FromBookAuthorEntity(bookId int, e *entity.BookAuthor) *BookAuthor {
	return &BookAuthor{
		BookId:   bookId,
		AuthorId: e.Id,
	}
}
