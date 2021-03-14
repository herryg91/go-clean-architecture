package crud_videos

import (
	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/entity"
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

func (uc *usecase) Get(id int) (*entity.Video, error) {
	return uc.repo.Get(id)
}

func (uc *usecase) GetAll() ([]*entity.Video, error) {
	return uc.repo.GetAll()
}

func (uc *usecase) Create(in entity.Video) (*entity.Video, error) {
	return uc.repo.Create(in)
}

func (uc *usecase) Update(in entity.Video) (*entity.Video, error) {
	return uc.repo.Update(in)
}

func (uc *usecase) Delete(id int) error {
	return uc.repo.Delete(id)
}
