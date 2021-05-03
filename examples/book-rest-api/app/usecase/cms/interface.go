package cms_usecase

import "github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"

type UseCase interface {
	GetBook(id int) (*entity.Book, error)
	GetBooks() ([]*entity.Book, error)
	CreateBook(in entity.Book) (id int, err error)

	GetAuthor(id int) (*entity.Author, error)
	GetAuthors() ([]*entity.Author, error)
	CreateAuthor(in entity.Author) (id int, err error)
}
