package crud_video

import (
	"errors"
	"fmt"

	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/app/repository"
	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/entity"
)

type usecase struct {
	video_repo repository.VideoRepository
}

func NewUseCase(video_repo repository.VideoRepository) UseCase {
	return &usecase{video_repo: video_repo}
}

func (uc *usecase) Get(id int) (*entity.Video, error) {
	data, err := uc.video_repo.Get(id)
	if err != nil {
		if errors.Is(err, repository.ErrVideoNotFound) {
			return nil, ErrVideoNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}

func (uc *usecase) GetAll() ([]*entity.Video, error) {
	data, err := uc.video_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}
func (uc *usecase) Create(in entity.Video) (*entity.Video, error) {
	data, err := uc.video_repo.Create(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}

func (uc *usecase) Update(in entity.Video) (*entity.Video, error) {
	data, err := uc.video_repo.Update(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil

}

func (uc *usecase) Delete(id int) error {
	err := uc.video_repo.Delete(id)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return err
}
