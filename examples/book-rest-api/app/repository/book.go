package repository

import (
	"errors"

	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"
)

var ErrBookNotFound = errors.New("Book not found")

type BookRepository interface {
	Get(id int) (*entity.Book, error)
	GetAll() ([]*entity.Book, error)
	Create(in entity.Book) (*entity.Book, error)
}
