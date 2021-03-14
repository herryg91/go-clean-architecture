package crud_books

import (
	"github.com/herryg91/go-clean-architecture/examples/simple-book-rest-api/entity"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) UseCase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Get(id int) (*entity.BookInfo, error) {
	return uc.repo.Get(id)
}

func (uc *usecase) GetAll() ([]*entity.BookInfo, error) {
	return uc.repo.GetAll()
}

func (uc *usecase) Create(in entity.Book, authorIds []int) error {
	return uc.repo.Create(in, authorIds)
}

func (uc *usecase) Update(in entity.Book, authorIds []int) error {
	return uc.repo.Update(in, authorIds)
}
