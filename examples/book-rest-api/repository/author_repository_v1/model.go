package author_repository_v1

import (
	"time"

	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/entity"
	"github.com/herryg91/go-clean-architecture/examples/book-rest-api/pkg/helpers"
)

type Author struct {
	Id        int          `gorm:"primary_key;column:id"`
	Name      string       `gorm:"column:name"`
	Birthdate helpers.Date `gorm:"column:birthdate"`
	CreatedAt *time.Time   `gorm:"column:created_at"`
	UpdatedAt *time.Time   `gorm:"column:updated_at"`
}

func (model *Author) ToAuthorEntity() *entity.Author {
	return &entity.Author{
		Id:        model.Id,
		Name:      model.Name,
		Birthdate: model.Birthdate,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
func (Author) FromAuthorEntity(in entity.Author) *Author {
	return &Author{
		Id:        in.Id,
		Name:      in.Name,
		Birthdate: in.Birthdate,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}
