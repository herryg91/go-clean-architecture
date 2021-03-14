package crud_books

import "github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/entity"

type Repository interface {
	Get(id int) (*entity.BookInfo, error)
	GetAll() ([]*entity.BookInfo, error)
	Create(in entity.Book, authorIds []int) error
	Update(in entity.Book, authorIds []int) error
}

type UseCase interface {
	Get(id int) (*entity.BookInfo, error)
	GetAll() ([]*entity.BookInfo, error)
	Create(in entity.Book, authorIds []int) error
	Update(in entity.Book, authorIds []int) error
}
