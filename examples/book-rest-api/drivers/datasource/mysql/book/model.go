package book_datasource

import (
	"time"

	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/pkg/helpers"
)

type BookModel struct {
	Id           int        `gorm:"primary_key;column:id"`
	Title        string     `gorm:"column:title"`
	ReleasedYear int        `gorm:"column:released_year"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
}

func (b *BookModel) ToBookEntity() *entity.Book {
	return &entity.Book{
		Id:           b.Id,
		Title:        b.Title,
		ReleasedYear: b.ReleasedYear,
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
	}
}

func (BookModel) FromBookEntity(book *entity.Book) *BookModel {
	return &BookModel{
		Id:           book.Id,
		Title:        book.Title,
		ReleasedYear: book.ReleasedYear,
		CreatedAt:    book.CreatedAt,
		UpdatedAt:    book.UpdatedAt,
	}
}

type BookAuthorJoinModel struct {
	BookModel
	AuthorId        int          `gorm:"column:author_id" json:"author_id"`
	AuthorName      string       `gorm:"column:author_name" json:"author_name"`
	AuthorBirthdate helpers.Date `gorm:"column:author_birthdate" json:"author_birthdate"`
}

func (m *BookAuthorJoinModel) ToEntity() (*entity.Book, *entity.Author) {
	book := m.ToBookEntity()
	author := &entity.Author{
		Id:        m.AuthorId,
		Name:      m.AuthorName,
		Birthdate: m.AuthorBirthdate,
	}
	return book, author
}
