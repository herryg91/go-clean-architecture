package crud_authors

import (
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"
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

func (uc *usecase) Get(id int) (*entity.Author, error) {
	return uc.repo.Get(id)
}

func (uc *usecase) GetAll() ([]*entity.Author, error) {
	return uc.repo.GetAll()
}

func (uc *usecase) Create(in entity.Author) (*entity.Author, error) {
	return uc.repo.Create(in)
}

func (uc *usecase) Update(in entity.Author) (*entity.Author, error) {
	return uc.repo.Update(in)
}
