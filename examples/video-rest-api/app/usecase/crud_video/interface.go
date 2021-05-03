package crud_video

import "github.com/herryg91/go-clean-architecture/examples/video-rest-api/entity"

type UseCase interface {
	Get(id int) (*entity.Video, error)
	GetAll() ([]*entity.Video, error)
	Create(in entity.Video) (*entity.Video, error)
	Update(in entity.Video) (*entity.Video, error)
	Delete(id int) error
}
