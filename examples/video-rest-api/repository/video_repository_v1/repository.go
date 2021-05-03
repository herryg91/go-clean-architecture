package video_repository_v1

import (
	"errors"
	"time"

	irepository "github.com/herryg91/go-clean-architecture/examples/video-rest-api/app/repository"
	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/entity"
	"gorm.io/gorm"
)

type repository struct {
	db        *gorm.DB
	tableName string
}

func New(db *gorm.DB) irepository.VideoRepository {
	return &repository{db, "videos"}
}

func (r *repository) Get(id int) (*entity.Video, error) {
	resp := VideoModel{}
	err := r.db.Table(r.tableName).Where("id = ?", id).First(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, irepository.ErrVideoNotFound
		}
		return nil, err
	}
	return resp.ToVideoEntity(), nil
}

func (r *repository) GetAll() ([]*entity.Video, error) {
	datas := []VideoModel{}
	err := r.db.Table(r.tableName).Find(&datas).Error
	if err != nil {
		return nil, err
	}
	resp := []*entity.Video{}

	for _, data := range datas {
		resp = append(resp, data.ToVideoEntity())
	}
	return resp, nil

}
func (r *repository) Create(in entity.Video) (*entity.Video, error) {
	videoModel := VideoModel{}.FromVideoEntity(in)

	timeNow := time.Now()
	videoModel.CreatedAt = &timeNow
	videoModel.UpdatedAt = &timeNow

	err := r.db.Table(r.tableName).Create(&videoModel).Error
	if err != nil {
		return nil, err
	}
	return videoModel.ToVideoEntity(), nil

}
func (r *repository) Update(in entity.Video) (*entity.Video, error) {
	videoModel := VideoModel{}.FromVideoEntity(in)
	timeNow := time.Now()
	videoModel.CreatedAt = nil
	videoModel.UpdatedAt = &timeNow
	err := r.db.Table(r.tableName).Where("id = ?", in.Id).Updates(&videoModel).Error
	if err != nil {
		return nil, err
	}
	return videoModel.ToVideoEntity(), nil

}
func (r *repository) Delete(id int) error {
	return r.db.Table(r.tableName).Delete(&VideoModel{}, "id = ?", id).Error
}
