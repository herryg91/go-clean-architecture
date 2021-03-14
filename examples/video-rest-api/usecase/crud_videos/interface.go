package crud_videos

import "github.com/herryg91/go-clean-architecture/examples/video-rest-api/entity"

type Repository interface {
	Get(id int) (*entity.Video, error)
	GetAll() ([]*entity.Video, error)
	Create(in entity.Video) (*entity.Video, error)
	Update(in entity.Video) (*entity.Video, error)
	Delete(id int) error
}

type UseCase interface {
	Get(id int) (*entity.Video, error)
	GetAll() ([]*entity.Video, error)
	Create(in entity.Video) (*entity.Video, error)
	Update(in entity.Video) (*entity.Video, error)
	Delete(id int) error
}
