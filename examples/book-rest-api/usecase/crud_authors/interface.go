package crud_authors

import "github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"

type Repository interface {
	Get(id int) (*entity.Author, error)
	GetAll() ([]*entity.Author, error)
	Create(in entity.Author) (*entity.Author, error)
	Update(in entity.Author) (*entity.Author, error)
}

type UseCase interface {
	Get(id int) (*entity.Author, error)
	GetAll() ([]*entity.Author, error)
	Create(in entity.Author) (*entity.Author, error)
	Update(in entity.Author) (*entity.Author, error)
}
