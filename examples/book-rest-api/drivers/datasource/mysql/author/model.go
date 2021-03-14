package author_datasource

import (
	"time"

	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/pkg/helpers"
)

type AuthorModel struct {
	Id        int          `gorm:"primary_key;column:id"`
	Name      string       `gorm:"column:name"`
	Birthdate helpers.Date `gorm:"column:birthdate"`
	CreatedAt *time.Time   `gorm:"column:created_at"`
	UpdatedAt *time.Time   `gorm:"column:updated_at"`
}

func (model *AuthorModel) ToAuthorEntity() *entity.Author {
	return &entity.Author{
		Id:        model.Id,
		Name:      model.Name,
		Birthdate: model.Birthdate,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
func (AuthorModel) FromAuthorEntity(in entity.Author) *AuthorModel {
	return &AuthorModel{
		Id:        in.Id,
		Name:      in.Name,
		Birthdate: in.Birthdate,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}

func (model *AuthorModel) ToBookAuthorEntity() *entity.BookAuthor {
	return &entity.BookAuthor{
		Id:   model.Id,
		Name: model.Name,
	}
}

type AuthorBookJoinModel struct {
	AuthorModel
	BookId           int    `gorm:"column:book_id"`
	BookTitle        string `gorm:"column:book_id"`
	BookReleasedYear int    `gorm:"column:book_released_year"`
}

func (m *AuthorBookJoinModel) ToEntity() (*entity.Author, *entity.Book) {
	author := m.ToAuthorEntity()
	book := &entity.Book{
		Id:           m.BookId,
		Title:        m.BookTitle,
		ReleasedYear: m.BookReleasedYear,
	}
	return author, book
}
