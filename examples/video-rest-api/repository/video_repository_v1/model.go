package video_repository_v1

import (
	"time"

	"github.com/herryg91/go-clean-architecture/examples/video-rest-api/entity"
)

type VideoModel struct {
	Id        int        `gorm:"primary_key;column:id"`
	Title     string     `gorm:"column:title"`
	Url       string     `gorm:"column:url"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (VideoModel) FromVideoEntity(v entity.Video) *VideoModel {
	return &VideoModel{
		Id:    v.Id,
		Title: v.Title,
		Url:   v.Url,
	}
}

func (m *VideoModel) ToVideoEntity() *entity.Video {
	return &entity.Video{
		Id:    m.Id,
		Title: m.Title,
		Url:   m.Url,
	}
}
