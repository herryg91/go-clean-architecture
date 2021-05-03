package repository

import (
	"errors"

	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"
)

var ErrAuthorNotFound = errors.New("Author not found")

type AuthorRepository interface {
	Get(id int) (*entity.Author, error)
	GetAll() ([]*entity.Author, error)
	Create(in entity.Author) (*entity.Author, error)
	Search(keyword string) ([]*entity.Author, error)
	MultiGet(ids []int) (map[int]*entity.Author, error)
}
